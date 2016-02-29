package KsanaDB

func getLuaScript(name string) string {
    name = name
    setTag :=` 
        local ret={};
        local tagHashName = ARGV[1];
        local seqArrayName = ARGV[2];
        local isKeyExist = 0;
        local seq = 0;
 
        for i,k in ipairs(KEYS) do
            isKeyExist = redis.call('HEXISTS', tagHashName, k);
            seq = 0;
            if isKeyExist == 0 then
                seq = redis.call('RPUSH', seqArrayName, k);
                redis.call('HSET', tagHashName, k, seq - 1);
            else
                seq = redis.call('HGET', tagHashName, k);
            end
            table.insert(ret, seq);
        end
        return ret;
    `
    return setTag
}
