学习笔记

1、递归go代码模版：
####
    func recur(level int, param int) {
        // terminator
        if level > MAX_LEVEL {
            // process result
            return
        }
        
        // process current logic
        process(level, param)
        
        // drill down
        recur(level + 1, newParam)
        
        // restore current status
    }
2、分治go代码模版
####
    func divide_conquer(problem *Problem, param int) int {
        // recursion terminator
        if problem == nil {
            process_result
            return return_result
        }
        
        // process current problem
        subproblems := split_problem(problem, data)
        subresult1 := divide_conquer(subproblem[0], p1)
        subresult2 := divide_conquer(subproblem[1], p1)
        subresult3 := divide_conquer(subproblem[2], p1)
        ...
        
        // merge
        result := process_result(subresult1, subresult2, subresult3)
        // revert the current level status
        
        return result
    }
3、回溯法：采用试错的思想，它尝试分步的去解决一个问题。在分步解决问题的过程中，当它通过尝试发现现有的分步答案不能得到有效的正确的解答的时候，它将取消上一步甚至是上几步的计算，再通过其它的可能的分步解答再次尝试寻找问题的答案。回溯法通常用最简单的递归方法来实现，在反复重复上述的步骤后可能出现两种情况：
####
- 找到一个可能存在的正确的答案；
- 在尝试了所有可能的分步方法后宣告该问题没有答案。
####
4、递归-循环，通过函数体自调用进行的循环
####
5、递归思维要点：
- 1、不要人肉进行递归（最大误区）
- 2、找到最近最简方法，将其拆解成可重复解决的问题（重复子问题）
- 3、数学归纳法思维