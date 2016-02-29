package KsanaDB

func getLuaScript(name string) string {
    setTag :=` 
        local ret={};
        local tagHashName = ARGV[1];
        local seqArrayName = ARGV[2];
        local isKeyExist = 0;
        local seq = -1;
 
        for i,k in ipairs(KEYS) do
            isKeyExist = redis.call('HEXISTS', tagHashName, k);
            seq = -1;
            if isKeyExist == 0 then
                seq = redis.call('RPUSH', seqArrayName, k);
                redis.call('HSET', tagHashName, k, seq - 1);
            end
            seq = redis.call('HGET', tagHashName, k);
            table.insert(ret, seq);
        end
        return cjson.encode(ret);
    `


    ret := "" 

    if name == "setTag" {
        ret = setTag
    }
    return ret
}
