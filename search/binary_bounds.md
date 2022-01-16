二分查找边界总结: ([参考](https://segmentfault.com/a/1190000016825704))
|查找方式|循环条件|左侧更新|右侧更新|中间点位置|返回值|
|:---:|:---:|:--:|:--:|:--:|:--:|
|标准二分查找|left <= right|left = mid - 1|right = mid + 1|(left + right) / 2	|-1 / mid|
|二分找左边界|left < right|left = mid - 1|right = mid|(left + right) / 2|	-1 / left|
|二分找右边界|left < right|left = mid|right = mid - 1|	(left + right) / 2 + 1|	-1 / right|