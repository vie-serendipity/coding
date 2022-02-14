import java.util.HashMap;

/*
 * @lc app=leetcode.cn id=208 lang=java
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
class Trie {

    class Node{
        public HashMap<Character, Node> table;
        public boolean isEnd;
        public Node() {
            isEnd=false;
            table = new HashMap<>();
        }
        
    }
    public Node root;
    public Trie() {
        root = new Node();
    }
    
    public void insert(String word) {
        Node p = root;
        
        for (int i = 0; i < word.length(); ++i) {
            if (p.table.containsKey(word.charAt(i))) {
                p = p.table.get(word.charAt(i));
            }
            else {
                p.table.put(word.charAt(i), new Node());
                p = p.table.get(word.charAt(i));
            }
        }
        p.isEnd=true;
        
    }
    
    public boolean search(String word) {
        Node p = root;
        int index = 0;
        while (p.table.containsKey(word.charAt(index))) {
            p = p.table.get(word.charAt(index));
            index++;
            if(index>=word.length()||p==null){
                if(p.isEnd==true){
                    return true;
                }
                else{
                    break;
                }
            }

        }
        return false;
    }
    
    public boolean startsWith(String prefix) {
        Node p = root;
        int index = 0;
        while (p.table.containsKey(prefix.charAt(index))) {
            p = p.table.get(prefix.charAt(index));
            index++;
            if(index==prefix.length()||p==null){
                return true;
            }
        }
        return false;
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * boolean param_2 = obj.search(word);
 * boolean param_3 = obj.startsWith(prefix);
 */
// @lc code=end

