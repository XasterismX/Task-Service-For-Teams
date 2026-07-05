package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"
)

type Limiter struct {
	tokens         float64
	maxTokens      float64
	refillRate     float64
	lastRefillTime time.Time
	mutex          sync.Mutex
}

func NewRateLimiter(maxTokens, refillRate float64) *Limiter {
	return &Limiter{
		tokens:         maxTokens,
		maxTokens:      maxTokens,
		refillRate:     refillRate,
		lastRefillTime: time.Now(),
	}
}

func RateLimitMiddleware(ipRateLimiter *IPRateLimiter) func(next http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				http.Error(w, "Invalid IP", http.StatusInternalServerError)
				return
			}

			limiter := ipRateLimiter.GetLimiter(ip)
			if limiter.Allow() {
				h.ServeHTTP(w, r)
			} else {
				http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
			}
		})
	}
}

func (r *Limiter) refillTokens() {
	now := time.Now()
	duration := now.Sub(r.lastRefillTime).Seconds()
	tokensToAdd := duration * r.refillRate

	r.tokens += tokensToAdd
	if r.tokens > r.maxTokens {
		r.tokens = r.maxTokens
	}
	r.lastRefillTime = now
}
func (r *Limiter) Allow() bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.refillTokens()

	if r.tokens >= 1 {
		r.tokens--
		return true
	}
	return false
}

type IPRateLimiter struct {
	limiters map[string]*Limiter
	mutex    sync.Mutex
}

func NewIPRateLimiter() *IPRateLimiter {
	return &IPRateLimiter{
		limiters: make(map[string]*Limiter),
	}
}

func (i *IPRateLimiter) GetLimiter(ip string) *Limiter {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	limiter, exists := i.limiters[ip]
	if !exists {
		// Allow 3 requests per minute
		limiter = NewRateLimiter(100, 0.05)
		i.limiters[ip] = limiter
	}

	return limiter
}
