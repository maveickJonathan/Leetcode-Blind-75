func isAnagram(s string, t string) bool {
    charSMap := make(map[byte]int)
    charTMap := make(map[byte]int)
    
    for i:=0; i < len(s); i++ {
        charSMap[s[i]] = charSMap[s[i]] + 1
    }
    
    for i:=0; i < len(t); i++ {
        charTMap[t[i]] = charTMap[t[i]] + 1
    }
    
    return reflect.DeepEqual(charSMap, charTMap)
}
