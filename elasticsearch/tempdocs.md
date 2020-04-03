GET /transactions/_search
{
    "from" : 0,
    "size": 1000,
    "query" : {
        "match_all" : {}
    }
}



GET /transactions/_search
{
    "query": {
        "multi_match" : {
            "query" : "AMZN",
            "fields" : ["description"]
        }
    }
}