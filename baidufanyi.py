
# -*- coding: UTF-8 -*-

import requests



def postTest(text): 
    print(text)
    url="https://dict.deepl.com/chinese-english/search?ajax=1&source=chinese&onlyDictEntries=1&translator=dnsof7h3k2lgh3gda&kind=full&eventkind=keyup&forleftside=true&il=zh"
    dataP={
        "query" : text.encoded(),
        "source" : "chinese",
        "onlyDictEntries" : "1" 
    }
    postD=requests.post(url,data=dataP)
    print(postD.url)
    print(postD.status_code)
    print(postD.headers)
    print(postD.text)

if __name__ == '__main__':
    postTest("你好")
