/*
 * @lc app=leetcode.cn id=51 lang=cpp
 *
 * [51] N 皇后
 */

// @lc code=start
#include<vector>
#include<string>
class Solution {
public:
    vector<vector<string>> solveNQueens(int n) {
        vector<vector<string>> res;
        vector<int> board;
        for(int i=0;i<n;++i){
            board.emplace_back(i);
            backtrack(board,n,res);
        }
        return res;
    }
    void backtrack(vector<int>& board,int n,vector<vector<string>>& res){
        if(board.size()==n){
            vector<string> tmp(n,string(n,'.'));
            for(int i=0;i<board.size();++i){
                tmp[i][board[i]]='Q';
            }
            res.emplace_back(tmp);
            board.pop_back();
            return;
        }
        for(int i=0;i<n;++i){
            if(check(board,i)){
                board.emplace_back(i);
                backtrack(board,n,res);
            }
        }
        board.pop_back();
    }
    bool check(vector<int> board,int i){
        for(int j=0;j<board.size();j++){
            if(i==board[j]||(j+board[j]==i+board.size())||(j-board[j]==board.size()-i)){
                return false;
            }
        }
        return true;
    }
};
// @lc code=end

