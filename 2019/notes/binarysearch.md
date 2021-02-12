## Binary Search
### Reference
[花花酱 LeetCode Binary Search - 刷题找工作 SP5](https://www.youtube.com/watch?v=v57lNF2mb_s)


+ Reduce the search space by half at each step
+ Input usually needs to be sorted

=>
+ 每一步都缩小一半的查找范围
+ 输入的数组需要有序

![201911091123342019-11-9-11-23-35.png](http://vscdn.pangzicaicai.com/undefined201911091123342019-11-9-11-23-35.png)

![201911091128482019-11-9-11-28-48.png](http://vscdn.pangzicaicai.com/undefined201911091128482019-11-9-11-28-48.png)

+ T(n) = T(n/2) + O(eval)
  + O(eval): 用于判定是否是解, 或者应该在搜索空间的左边还是右边

![201911091133122019-11-9-11-33-13.png](http://vscdn.pangzicaicai.com/undefined201911091133122019-11-9-11-33-13.png)

```python
# search [l, r)
def binary_search(l, r):
    while l < r:
        m = l + (r - l) // 2 # 求算中点
        if f(m): return m # optional
        if g(m): 
            r = m #new range [l, m)
        else:
            l = m + 1 # new range [m + 1, r]
    return l # or not found 

# search[l, r]
def binary_search(l, r):
    while l <= r:
        m = l + (r - l) // 2
        if f(m): return m # optional
        if g(m):
            r = m - 1 # new range [l, m - 1]
        else:
            l = m + 1 # new range [m + 1, r]
    return l # or not found
```
+ 为什么最后返回的是l, 而不是r:
  + 在while中, 最后一次是 l == r, 则有 l=r=mid
  + 当 list[mid] >= target, r = mid - 1, r = l - 1
  + 当 list[mid] < target, l = mid + 1, l = r + 1
  + l均在r的右侧1位即是, l = r + 1,
  + 题要求最小的大于等于target的idx
    + 当 最后一次的 m 有 list[m] >= target 时, l的位置不变
    + 当 最后一次 m 有 list[m] < target时, l 刚好 + 1可以大于 target

![201911091146522019-11-9-11-46-53.png](http://vscdn.pangzicaicai.com/undefined201911091146522019-11-9-11-46-53.png)

针对于可以存在重复元素的有序数组, 为了可以找到某个元素的起始下标和终点下标有如下两个方法:
+ lower_bound: 找到起始下标(闭), 即是第一个 >= target 的 index
(找第一个 >= target的 index,)
+ upper_bound:
终点下标(开), 即是第一个 > target 的 index
(找第一个 > target的index,最后一个 == target的index的下一个)

![201911091159432019-11-9-11-59-43.png](http://vscdn.pangzicaicai.com/undefined201911091159432019-11-9-11-59-43.png)

+ lc 69.   
sqrt(x)
希望找到一个最小的数的平方, 是的这个数的平方 > target

+ lc. 70.  
即是找到`最小的`那个badVersion

![201911091205472019-11-9-12-5-48.png](http://vscdn.pangzicaicai.com/undefined201911091205472019-11-9-12-5-48.png)

+ 378. 找到 k 个 小于这个数的数值

![201911091308332019-11-9-13-8-33.png](http://vscdn.pangzicaicai.com/undefined201911091308332019-11-9-13-8-33.png)

![201911091314072019-11-9-13-14-10.png](http://vscdn.pangzicaicai.com/undefined201911091314072019-11-9-13-14-10.png)
右边的返回的是r+1

![201911091355052019-11-9-13-55-6.png](http://vscdn.pangzicaicai.com/undefined201911091355052019-11-9-13-55-6.png)

![201911091356382019-11-9-13-56-39.png](http://vscdn.pangzicaicai.com/undefined201911091356382019-11-9-13-56-39.png)