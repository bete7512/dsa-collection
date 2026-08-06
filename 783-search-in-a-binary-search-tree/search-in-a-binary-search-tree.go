/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil{
        return nil
    }


    cur := root
    for {
        if val == cur.Val{
           return cur
        }

        if val < cur.Val{
            if cur.Left != nil{
                cur = cur.Left
                continue
            }
            return nil
        } else {
             if cur.Right != nil{
                cur = cur.Right
                continue
            }
            return nil
        }
    }
}