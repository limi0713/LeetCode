/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU缓存
 *
 * https://leetcode-cn.com/problems/lfu-cache/description/
 *
 * algorithms
 * Hard (34.36%)
 * Likes:    112
 * Dislikes: 0
 * Total Accepted:    4.3K
 * Total Submissions: 11.7K
 * Testcase Example:  '["LFUCache","put","put","get","put","get","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]'
 *
 * 设计并实现最不经常使用（LFU）缓存的数据结构。它应该支持以下操作：get 和 put。
 * 
 * get(key) - 如果键存在于缓存中，则获取键的值（总是正数），否则返回 -1。
 * put(key, value) -
 * 如果键不存在，请设置或插入值。当缓存达到其容量时，它应该在插入新项目之前，使最不经常使用的项目无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，最近最少使用的键将被去除。
 * 
 * 进阶：
 * 你是否可以在 O(1) 时间复杂度内执行两项操作？
 * 
 * 示例：
 * 
 * 
 * LFUCache cache = new LFUCache( 2 /* capacity (缓存容量)
 * 
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回 1
 * cache.put(3, 3);    // 去除 key 2
 * cache.get(2);       // 返回 -1 (未找到key 2)
 * cache.get(3);       // 返回 3
 * cache.put(4, 4);    // 去除 key 1
 * cache.get(1);       // 返回 -1 (未找到 key 1)
 * cache.get(3);       // 返回 3
 * cache.get(4);       // 返回 4
 * 
 */
// @lc code=start
type LFUCache struct {
    len int
    capacity int
    nodes map[int]*FreqNode
    lsFreq *FreqNode
}

type cacheNode struct{
    key int
    value int
}

type FreqNode struct{
    cache *cacheNode
    freq int
    preNode *FreqNode
    nextNode *FreqNode
}

func Constructor(capacity int) LFUCache {

    return LFUCache{
        capacity:capacity,
        nodes : make(map[int]*FreqNode,capacity),
        lsFreq : &FreqNode{},
    }
}

func (this *LFUCache) Get(key int) int {

    if v,ok := this.nodes[key]; ok {
        v.freq += 1
        this.adjustOrder(v)
        return v.cache.value
    }

    return -1
}

func (this *LFUCache)adjustOrder(node *FreqNode){
    for next := node.nextNode;next!=nil;{
        if next.freq > node.freq {
            break
        }
        pre := node.preNode
        pre.nextNode = next
        node.nextNode = next.nextNode
        next.nextNode = node
        next.preNode = pre
        node.preNode = next 
        next = node.nextNode
        if next != nil {
            node.nextNode.preNode = node
        }
    }
}

func (this *LFUCache) Put(key int, value int)  {
    if this.capacity <= 0 {
        return
    }

    if v,ok := this.nodes[key];ok {
        v.cache.value = value
        v.freq += 1
        this.adjustOrder(v)
        return
    }

    if this.len < this.capacity {
        node := &FreqNode{
            cache:&cacheNode{
                key:key,
                value:value,
            },
            freq:1,
            nextNode:this.lsFreq.nextNode,
            preNode : this.lsFreq,
        }

        this.lsFreq.nextNode = node
        if node.nextNode != nil {
            node.nextNode.preNode = node
        }

        this.adjustOrder(node)

        this.nodes[key] = node

        this.len ++ 
        return
    }

    if this.len == this.capacity {
        rmNode := this.lsFreq.nextNode
    
        this.lsFreq.nextNode = rmNode.nextNode
        if rmNode.nextNode != nil {
            rmNode.nextNode.preNode = this.lsFreq
        }
        delete(this.nodes,rmNode.cache.key)

        this.len -- 

        node := &FreqNode{
            cache:&cacheNode{
                key:key,
                value:value,
            },
            freq:1,
            nextNode:this.lsFreq.nextNode,
            preNode : this.lsFreq,
        }

        this.lsFreq.nextNode = node
        if node.nextNode != nil {
            node.nextNode.preNode = node
        }

        this.adjustOrder(node)

        this.nodes[key] = node

        this.len ++ 
    }
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

