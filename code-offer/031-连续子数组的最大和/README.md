[toc]

## 题目
- [题目](https://blog.csdn.net/gatieme/article/details/51287801)

## 题目描述
```text
HZ偶尔会拿些专业问题来忽悠那些非计算机专业的同学。
今天测试组开完会后,他又发话了:在古老的一维模式识别中,常常需要计算连续子向量的最大和,
当向量全为正数的时候,问题很好解决。
但是,如果向量中包含负数,是否应该包含某个负数,并期望旁边的正数会弥补它呢？
例如:{6,-3,-2,7,-15,1,2,2},连续子向量的最大和为8(从第0个开始,到第3个为止)。
给一个数组，返回它的最大连续子序列的和，你会不会被他忽悠住？(子向量的长度至少是1)
```

## 样例
### 输入
```text
6,-3,-2,7,-15,1,2,2
```

### 输出
```text
0,3
```

## 思路
### 方法一：暴力方法
最直接的方法就是找出所有的子数组，然后求其和，取最大

### 方法二：动态规划


### 方法三：贪心思想
