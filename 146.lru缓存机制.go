/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 *
 * https://leetcode-cn.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (46.67%)
 * Likes:    478
 * Dislikes: 0
 * Total Accepted:    43.4K
 * Total Submissions: 92.9K
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 * 
 * 获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
 * 写入数据 put(key, value) -
 * 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
 * 
 * 进阶:
 * 
 * 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 * 
 * 示例:
 * 
 * LRUCache cache = new LRUCache( 2 /* 缓存容量 */ /*);
 * 
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);       // 返回  1
 * cache.put(3, 3);    // 该操作会使得密钥 2 作废
 * cache.get(2);       // 返回 -1 (未找到)
 * cache.put(4, 4);    // 该操作会使得密钥 1 作废
 * cache.get(1);       // 返回 -1 (未找到)
 * cache.get(3);       // 返回  3
 * cache.get(4);       // 返回  4
 * 
 * 
 */

// @lc code=start

type cacheNode struct{
    key int
    value int
}

type useNode struct{
    cache *cacheNode
    preNode *useNode
    nextNode *useNode
}

type LRUCache struct {
    len int
    capacity int
    nodes map[int]*useNode
    lsUseHead *useNode
    lsUseEnd *useNode
}


func Constructor(capacity int) LRUCache {
    lc := LRUCache{
        len : 0,
        capacity : capacity,
        nodes : make(map[int]*useNode,capacity),
        lsUseHead : &useNode{},
        lsUseEnd : &useNode{},
    }

    lc.lsUseHead.nextNode = lc.lsUseEnd
    lc.lsUseEnd.preNode = lc.lsUseHead
    return lc
}

func (this *LRUCache)adjustOrder(node *useNode){
    pre := node.preNode
    next := node.nextNode

    pre.nextNode = next
    next.preNode = pre

    node.nextNode = this.lsUseEnd
    node.preNode = this.lsUseEnd.preNode
    node.preNode.nextNode = node
    this.lsUseEnd.preNode = node
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.nodes[key]; ok {
        this.adjustOrder(node)
        return node.cache.value
    }

    return -1
}

func (this *LRUCache) Put(key int, value int)  {
    if this.capacity == 0 {
        return
    }

    if node, ok := this.nodes[key]; ok {
        node.cache.value = value
        this.adjustOrder(node)
        return
    }

    if this.len < this.capacity {
        node := &useNode{
            cache : &cacheNode{
                key : key,
                value : value,
            },
            preNode : this.lsUseEnd.preNode,
            nextNode : this.lsUseEnd,
        }

        this.lsUseEnd.preNode.nextNode = node
        this.lsUseEnd.preNode = node
        this.nodes[key] = node
        this.len ++
        return
    }
    rm := this.lsUseHead.nextNode
    this.lsUseHead.nextNode = rm.nextNode
    rm.nextNode.preNode = this.lsUseHead
    delete(this.nodes,rm.cache.key)
    this.len -- 

    node := &useNode{
        cache : &cacheNode{
            key : key,
            value : value,
        },
        preNode : this.lsUseEnd.preNode,
        nextNode : this.lsUseEnd,
    }

    this.lsUseEnd.preNode.nextNode = node
    this.lsUseEnd.preNode = node
    this.nodes[key] = node
    this.len ++
    return
    
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

