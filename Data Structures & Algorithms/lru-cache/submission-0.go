type LRUCache struct {
    cache        map[int]int
    usage        map[int]int
    capacity     int
    current_fill int
    timer        int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        cache:        make(map[int]int),
        usage:        make(map[int]int),
        capacity:     capacity,
        current_fill: 0,
        timer:        0,
    }
}

func (this *LRUCache) Get(key int) int {
    if _, exists := this.cache[key]; exists {
        this.timer += 1          // Time moves forward
        this.usage[key] = this.timer // Update to the absolute latest timestamp
        return this.cache[key]
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    this.timer += 1 // Every operation moves time forward

    if _, exists := this.cache[key]; exists {
        this.cache[key] = value
        this.usage[key] = this.timer // Update timestamp
    } else {
        if this.current_fill == this.capacity {
            // Find the oldest timestamped key and remove it
            oldest_time := 10000000 // A massive number
            min_key := 0
            for k, timestamp := range this.usage {
                if timestamp < oldest_time {
                    oldest_time = timestamp
                    min_key = k
                }
            }
            delete(this.cache, min_key)
            delete(this.usage, min_key)
            this.current_fill -= 1 // Vacate the slot
        }
        
        this.cache[key] = value
        this.usage[key] = this.timer
        this.current_fill += 1
    }
}