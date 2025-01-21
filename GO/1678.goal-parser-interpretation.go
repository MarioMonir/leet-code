/*
 * @lc app=leetcode id=1678 lang=golang
 *
 * [1678] Goal Parser Interpretation
 */

package main


// @lc code=start
func interpret(command string) string {
    res := ""
    for i:=0 ; i < len(command); i++ {
        if string(command[i]) == "G" {
            res += "G"
        }else if i < len(command)-1 && string(command[i:i+2]) == "()" {
            res += "o"
            i+=1
        }else if i < len(command)-3 && string(command[i:i+4])== "(al)" {
            res += "al"
            i+=3
        }
    }  
    return res 
}
// @lc code=end


