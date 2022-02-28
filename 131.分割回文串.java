/*
 * @lc app=leetcode.cn id=131 lang=java
 *
 * [131] 分割回文串
 */

// @lc code=start
class Solution {
    boolean[][] f;
    public List<List<String>> partition(String s) {
        int n=s.length();
        f=new boolean[n][n];
        for(int i=0;i<n;++i){
            Arrays.fill(f[i],true);
        }
        for(int i=n-1;i>=0;--i){
            for(int j=i+1;j<n;++j){
                f[i][j]=(s.charAt(i)==s.charAt(j))&&f[i+1][j-1];
            }
        }
        List<List<String>> res=new ArrayList<>();
        List<String> temp=new ArrayList<>();
        backtrack(s,0,s.length(),temp,res);
        return res;
    }
    private void backtrack(String s,int start,int end,List<String> temp,List<List<String>> res){
        if(start==end){
            List<String> t=new ArrayList<>(temp);
            res.add(t);
            return;
        }
        for(int i=start+1;i<=end;++i){
            if(f[start][i-1]){
                temp.add(s.substring(start,i));
                backtrack(s,i,end,temp,res);
                temp.remove(temp.size()-1);
            }
        }
    }
}
// @lc code=end

