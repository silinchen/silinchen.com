# Flex 布局



### 一、Flex 介绍

Flex 全称 Flexible box 布局模型，通常称为  flexbox 或 flex。一种比较高效率的 css3 布局方案。Flex 有两根轴线，分别是主轴和交叉轴。 

flex 是一种一维的布局，是因为一个 flex 一次只能处理一个维度上的元素布局，一行（左右）或者一列（上下），通过主轴控制子元素排列的方向（上下/左右）,交叉轴始终垂直于主轴，配合排列。作为对比的是另外一个二维布局 [CSS Grid Layout](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Grid_Layout)，可以同时处理行和列上的布局。 



### 二、相关参数

与 Flex 有关的参数，主要包括一下几个：

1. `dispaly` - 参数值设为 flex 或 inline-flex 即指定布局方式为 flex，指定后子元素的 `float`、`clear`和 `vertical-align` 属性失效。
2.  `flex-direction` - 参数用于指定主轴的方向，有4个值  `row` 、`row-reverse`、 `colunm` 、`colunm-reverse` 。`row` 是行（与 `inline` 行为一致，从左向右 → ），`colunm` 是列（与 `block` 行为一致，从上到下 `↓` ）。带上 `reverse` 的是指反向，即原来 `row` 是左右 `→` ，`row-reverse` 就是右左 `←`，`column` 是上下 `↓`，则 `column` 就是下上 `↑` 。
3. 