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
        local function compare(a,b)
            return a<b;
        end
        table.sort(ret,compare);
        return cjson.encode(ret);
    `

    // http://lua-users.org/wiki/SplitJoin
    getTag := `
        local function split (str, pat) 
            local t = {};
            local fpat = "(.-)" .. pat;
            local last_end = 1;
            local s, e, cap = str:find(fpat, 1);
            while s do 
                if s ~= 1 or cap ~= "" then
                    table.insert(t, cap);
                end
                last_end = e + 1;
                s, e, cap = str:find(fpat, last_end);
            end
            if last_end <= #str then
                cap = str:sub(last_end);
                table.insert(t, cap);
            end
            return t;
        end

        local function all (tagSeq)
            local ret = {}; 
            local kv = {};
            for i,v in ipairs(tagSeq) do
                kv = split(v, "\t");
                if ret[kv[1]] == nil then
                    ret[kv[1]] = {};
                end
                table.insert(ret[kv[1]], kv[2]); 
            end
            return ret;
        end

        local function tag(tagSeq)
            local ret = {}; 
            local kv = {};
            for i,v in ipairs(tagSeq) do
                kv = split(v, "\t");
                ret[#ret+1] = kv[1]; 
            end
            return ret;
        end

        local seqArrayName = ARGV[1];
        local target = ARGV[2];
        local tagName = ARGV[3]; -- only for get tag content
        local tagSeq = {}; 
        local ret = {}; 

        tagSeq = redis.call("LRANGE" ,seqArrayName, 0, -1);

        if target == "All" then
            ret = all(tagSeq);
        elseif target == "TagKey" then
            ret = tag(tagSeq);
        elseif target == "TagValue" then
            local tmp = all(tagSeq);
            ret = tmp[tagName];
        end
        return cjson.encode(ret);
    `

    ret := "" 

    if name == "setTag" {
        ret = setTag
    } else if name == "getTag" {
        ret = getTag    
    }
    return ret
}
