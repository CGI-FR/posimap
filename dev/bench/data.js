window.BENCHMARK_DATA = {
  "lastUpdate": 1747228204885,
  "repoUrl": "https://github.com/CGI-FR/posimap",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "adrien.aury@cgi.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "adrien.aury@cgi.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "distinct": true,
          "id": "24dfdcf697e12373e9ead1f860aefc2e5dfc776c",
          "message": "chore: add benchmark tests",
          "timestamp": "2025-04-30T15:34:38Z",
          "tree_id": "dbad118df31fcd550b432b18175bd03f0fdb237c",
          "url": "https://github.com/CGI-FR/posimap/commit/24dfdcf697e12373e9ead1f860aefc2e5dfc776c"
        },
        "date": 1746028201050,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 1931450,
            "unit": "ns/op\t 1700737 B/op\t    3446 allocs/op",
            "extra": "6109 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 1931450,
            "unit": "ns/op",
            "extra": "6109 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 1700737,
            "unit": "B/op",
            "extra": "6109 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 3446,
            "unit": "allocs/op",
            "extra": "6109 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold",
            "value": 1912957,
            "unit": "ns/op\t 1700516 B/op\t    3443 allocs/op",
            "extra": "6265 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - ns/op",
            "value": 1912957,
            "unit": "ns/op",
            "extra": "6265 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - B/op",
            "value": 1700516,
            "unit": "B/op",
            "extra": "6265 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - allocs/op",
            "value": 3443,
            "unit": "allocs/op",
            "extra": "6265 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "adrien.aury@cgi.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "adrien.aury@cgi.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "distinct": true,
          "id": "8b2d1dde7b219dbb06d8f78dbd0933c5ef0263c8",
          "message": "docs: complete changelog",
          "timestamp": "2025-04-30T15:59:47Z",
          "tree_id": "f193a4c08bf64e5364b8306240e494388ffe1625",
          "url": "https://github.com/CGI-FR/posimap/commit/8b2d1dde7b219dbb06d8f78dbd0933c5ef0263c8"
        },
        "date": 1746028912416,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 1911730,
            "unit": "ns/op\t 1700730 B/op\t    3446 allocs/op",
            "extra": "6218 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 1911730,
            "unit": "ns/op",
            "extra": "6218 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 1700730,
            "unit": "B/op",
            "extra": "6218 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 3446,
            "unit": "allocs/op",
            "extra": "6218 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold",
            "value": 1922079,
            "unit": "ns/op\t 1700529 B/op\t    3443 allocs/op",
            "extra": "6300 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - ns/op",
            "value": 1922079,
            "unit": "ns/op",
            "extra": "6300 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - B/op",
            "value": 1700529,
            "unit": "B/op",
            "extra": "6300 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - allocs/op",
            "value": 3443,
            "unit": "allocs/op",
            "extra": "6300 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "d4991f7dcb87cfa2a37ca52e3eb0ff21b6128273",
          "message": "refactor: fixed-width encoding (#1)\n\n* refactor: wip! buffer and string codec\n\n* refactor: wip! record value\n\n* refactor: wip! record value\n\n* refactor: wip! record object\n\n* refactor: wip! rename method\n\n* fix: wip! remove unused error\n\n* fix: wip! export is responsability of object\n\n* fix: wip! introduce named record\n\n* fix: wip! new approach for records\n\n* fix: wip! export/import objects\n\n* fix: wip! primitive + remove named and rewrite object\n\n* fix: wip! export feedback idea\n\n* fix: wip! here it is the genius idea\n\n* refactor: wip! test ok with object\n\n* refactor: wip! hide implementation detail for record\n\n* refactor: wip! add array record\n\n* refactor: wip! externalize struct tokens in its own package\n\n* style: lint\n\n* refactor: rewrite/rename struct to document\n\n* refactor: remove unused api\n\n* refactor: wip! predicate\n\n* refactor: predicate\n\n* refactor: wip! remove useless tokens\n\n* refactor: import\n\n* refactor: wip! buffers\n\n* refactor: wip! jsonline writer\n\n* refactor: wip! fix buffer and jsonline writer\n\n* refactor: wip! write spaces on trimmed strings\n\n* refactor: wip! feat schema\n\n* refactor: wip! buffer read next\n\n* refactor: wip! rename Memory to Buffer\n\n* refactor: move packages\n\n* refactor: wip! config\n\n* refactor: wip! config\n\n* refactor: wip! compile config\n\n* refactor: wip! fold command\n\n* refactor: wip! fix compile config\n\n* refactor: wip! fix writer flush\n\n* refactor: write last nl\n\n* refactor: wip! fold command trim\n\n* refactor: unfold command\n\n* test: benchmark discard output",
          "timestamp": "2025-05-07T21:43:11+02:00",
          "tree_id": "ee967873d900500ee6f491aa72eac9b5b26def20",
          "url": "https://github.com/CGI-FR/posimap/commit/d4991f7dcb87cfa2a37ca52e3eb0ff21b6128273"
        },
        "date": 1746647112684,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 962106,
            "unit": "ns/op\t  781671 B/op\t    3809 allocs/op",
            "extra": "12463 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 962106,
            "unit": "ns/op",
            "extra": "12463 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 781671,
            "unit": "B/op",
            "extra": "12463 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 3809,
            "unit": "allocs/op",
            "extra": "12463 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold",
            "value": 920801,
            "unit": "ns/op\t  625560 B/op\t    5688 allocs/op",
            "extra": "13035 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - ns/op",
            "value": 920801,
            "unit": "ns/op",
            "extra": "13035 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - B/op",
            "value": 625560,
            "unit": "B/op",
            "extra": "13035 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - allocs/op",
            "value": 5688,
            "unit": "allocs/op",
            "extra": "13035 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "02e63fe66116c617d742143dd738125268009c2d",
          "message": "feat: configure default encoding via flag (#20)\n\n* chore: fix fixed-width file editor config\n\n* test: fix test data schema definition\n\n* fix: do not trim control characters\n\n* fix: json string encoding\n\n* refactor: config is on the appli side, not infra\n\n* fix: flag config should not override yaml config\n\n* fix: flag config should not override yaml config\n\n* feat: configure default encoding via flag",
          "timestamp": "2025-05-09T14:00:33+02:00",
          "tree_id": "26d81ac63047b2dde333f6e5d348ef1f5cc6c551",
          "url": "https://github.com/CGI-FR/posimap/commit/02e63fe66116c617d742143dd738125268009c2d"
        },
        "date": 1746792138028,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 1357747,
            "unit": "ns/op\t 1054970 B/op\t    6588 allocs/op",
            "extra": "8523 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 1357747,
            "unit": "ns/op",
            "extra": "8523 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 1054970,
            "unit": "B/op",
            "extra": "8523 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 6588,
            "unit": "allocs/op",
            "extra": "8523 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "d6768de6b54bd49f2fcb14cdaa4f4cae63f278ab",
          "message": "feat: configure record length in yaml config (#22)",
          "timestamp": "2025-05-09T16:00:42+02:00",
          "tree_id": "a2e908380cbf92b4ee833194152bdadd47236364",
          "url": "https://github.com/CGI-FR/posimap/commit/d6768de6b54bd49f2fcb14cdaa4f4cae63f278ab"
        },
        "date": 1746799368879,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 1413069,
            "unit": "ns/op\t 1109167 B/op\t    6864 allocs/op",
            "extra": "8426 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 1413069,
            "unit": "ns/op",
            "extra": "8426 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 1109167,
            "unit": "B/op",
            "extra": "8426 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 6864,
            "unit": "allocs/op",
            "extra": "8426 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold",
            "value": 942286,
            "unit": "ns/op\t  629816 B/op\t    5777 allocs/op",
            "extra": "12753 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - ns/op",
            "value": 942286,
            "unit": "ns/op",
            "extra": "12753 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - B/op",
            "value": 629816,
            "unit": "B/op",
            "extra": "12753 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - allocs/op",
            "value": 5777,
            "unit": "allocs/op",
            "extra": "12753 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "adrien.aury@cgi.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "adrien.aury@cgi.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "distinct": true,
          "id": "b7dfc489dae38a3b63bd8a77223f5c2adf6167f5",
          "message": "fix: unfold clear buffer and marshalers",
          "timestamp": "2025-05-09T21:24:33Z",
          "tree_id": "d92bfbb740bf66198049c0c60aca2af4e20c6307",
          "url": "https://github.com/CGI-FR/posimap/commit/b7dfc489dae38a3b63bd8a77223f5c2adf6167f5"
        },
        "date": 1746826004015,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 1382247,
            "unit": "ns/op\t 1109162 B/op\t    6864 allocs/op",
            "extra": "8401 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 1382247,
            "unit": "ns/op",
            "extra": "8401 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 1109162,
            "unit": "B/op",
            "extra": "8401 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 6864,
            "unit": "allocs/op",
            "extra": "8401 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold",
            "value": 930212,
            "unit": "ns/op\t  627526 B/op\t    5054 allocs/op",
            "extra": "12894 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - ns/op",
            "value": 930212,
            "unit": "ns/op",
            "extra": "12894 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - B/op",
            "value": 627526,
            "unit": "B/op",
            "extra": "12894 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - allocs/op",
            "value": 5054,
            "unit": "allocs/op",
            "extra": "12894 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "distinct": true,
          "id": "7b5e908c145a86d584ac4d7b53229965908f086d",
          "message": "refactor: use generics for codec interfaces",
          "timestamp": "2025-05-11T07:49:10Z",
          "tree_id": "08bc80c391046e7009ba1cdf6ea5d1ee21c87cb0",
          "url": "https://github.com/CGI-FR/posimap/commit/7b5e908c145a86d584ac4d7b53229965908f086d"
        },
        "date": 1746949868526,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 1389635,
            "unit": "ns/op\t 1108501 B/op\t    6864 allocs/op",
            "extra": "8329 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 1389635,
            "unit": "ns/op",
            "extra": "8329 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 1108501,
            "unit": "B/op",
            "extra": "8329 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 6864,
            "unit": "allocs/op",
            "extra": "8329 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold",
            "value": 935569,
            "unit": "ns/op\t  626870 B/op\t    5054 allocs/op",
            "extra": "12825 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - ns/op",
            "value": 935569,
            "unit": "ns/op",
            "extra": "12825 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - B/op",
            "value": 626870,
            "unit": "B/op",
            "extra": "12825 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - allocs/op",
            "value": 5054,
            "unit": "allocs/op",
            "extra": "12825 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "2923eae972512ac8445a84c919c8a286725f8a60",
          "message": "feat: schema graph and auto-add fillers (#23)\n\n* refactor: wip! schema graph modelisation\n\n* refactor: wip! compile marshaling path\n\n* refactor: wip! compute offsets\n\n* refactor: wip! compute sizes\n\n* refactor: wip! take redefines into account\n\n* refactor: wip! fix size compute\n\n* refactor: wip! test expected output\n\n* refactor: wip! fix missing fillers\n\n* refactor: wip! fix missing fillers\n\n* refactor: wip! option to show dependsOn\n\n* refactor: wip! node use unique ids\n\n* refactor: wip! added schema validation\n\n* refactor: wip! log errors on validation\n\n* refactor: wip! take into account occurs\n\n* refactor: wip! use refactored code and add graph command\n\n* fix: add missing fillers\n\n* refactor(schema): remove old impl and fix tests\n\n* refactor(schema): final rename",
          "timestamp": "2025-05-12T10:49:02+02:00",
          "tree_id": "90c0fc2107af9ce8e0b40b468efbf0c9b9522433",
          "url": "https://github.com/CGI-FR/posimap/commit/2923eae972512ac8445a84c919c8a286725f8a60"
        },
        "date": 1747039859404,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 1407604,
            "unit": "ns/op\t 1127607 B/op\t    7045 allocs/op",
            "extra": "8341 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 1407604,
            "unit": "ns/op",
            "extra": "8341 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 1127607,
            "unit": "B/op",
            "extra": "8341 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 7045,
            "unit": "allocs/op",
            "extra": "8341 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold",
            "value": 940670,
            "unit": "ns/op\t  635883 B/op\t    5217 allocs/op",
            "extra": "12764 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - ns/op",
            "value": 940670,
            "unit": "ns/op",
            "extra": "12764 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - B/op",
            "value": 635883,
            "unit": "B/op",
            "extra": "12764 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - allocs/op",
            "value": 5217,
            "unit": "allocs/op",
            "extra": "12764 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "37492cdbbd1b33b0a7b0ff6e2cf935acfb060d1f",
          "message": "fix: import document any order (#27)\n\n* fix: import document any order\n\n* test: fix record import unit test\n\n* fix: import array out of bound exception\n\n* chore: do not add final new line to fixed-width files",
          "timestamp": "2025-05-12T13:23:12+02:00",
          "tree_id": "3cf1c70a811469be2f618ffe386925ad19c6e872",
          "url": "https://github.com/CGI-FR/posimap/commit/37492cdbbd1b33b0a7b0ff6e2cf935acfb060d1f"
        },
        "date": 1747049114346,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 1416320,
            "unit": "ns/op\t 1127606 B/op\t    7045 allocs/op",
            "extra": "8202 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 1416320,
            "unit": "ns/op",
            "extra": "8202 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 1127606,
            "unit": "B/op",
            "extra": "8202 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 7045,
            "unit": "allocs/op",
            "extra": "8202 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold",
            "value": 877284,
            "unit": "ns/op\t  682145 B/op\t    3707 allocs/op",
            "extra": "13687 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - ns/op",
            "value": 877284,
            "unit": "ns/op",
            "extra": "13687 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - B/op",
            "value": 682145,
            "unit": "B/op",
            "extra": "13687 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - allocs/op",
            "value": 3707,
            "unit": "allocs/op",
            "extra": "13687 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "2bcb80d304734722655627a13ec6445211ee267e",
          "message": "feat: rename config flag to schema (#29)\n\n* feat: rename config flag to schema\n\n* style: showDenpendencies should be showDependencies\n\n* test: rename config flag to schema",
          "timestamp": "2025-05-12T13:59:56+02:00",
          "tree_id": "6044568d4a4bac39dfd66b7c637deafc42d263d1",
          "url": "https://github.com/CGI-FR/posimap/commit/2bcb80d304734722655627a13ec6445211ee267e"
        },
        "date": 1747051313067,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 1418620,
            "unit": "ns/op\t 1127603 B/op\t    7045 allocs/op",
            "extra": "8260 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 1418620,
            "unit": "ns/op",
            "extra": "8260 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 1127603,
            "unit": "B/op",
            "extra": "8260 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 7045,
            "unit": "allocs/op",
            "extra": "8260 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold",
            "value": 878815,
            "unit": "ns/op\t  682157 B/op\t    3707 allocs/op",
            "extra": "13659 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - ns/op",
            "value": 878815,
            "unit": "ns/op",
            "extra": "13659 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - B/op",
            "value": 682157,
            "unit": "B/op",
            "extra": "13659 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - allocs/op",
            "value": 3707,
            "unit": "allocs/op",
            "extra": "13659 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "60bc36d34b4aff6ef025b6b2d7518b0f4fcc862c",
          "message": "feat: add charsets command (#30)",
          "timestamp": "2025-05-12T16:15:43+02:00",
          "tree_id": "f451b3c9d434259654c0ceed6617027bffaed4c0",
          "url": "https://github.com/CGI-FR/posimap/commit/60bc36d34b4aff6ef025b6b2d7518b0f4fcc862c"
        },
        "date": 1747059446522,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 1437740,
            "unit": "ns/op\t 1127601 B/op\t    7045 allocs/op",
            "extra": "7554 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 1437740,
            "unit": "ns/op",
            "extra": "7554 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 1127601,
            "unit": "B/op",
            "extra": "7554 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 7045,
            "unit": "allocs/op",
            "extra": "7554 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "9f283e822e2cd0c26dce02b7ccbe504ec6dbd3ad",
          "message": "feat: configure record feedback from yaml (#31)\n\n* feat: configure record feedback from yaml\n\n* test: adapt bench tests\n\n* style: lint",
          "timestamp": "2025-05-12T17:19:09+02:00",
          "tree_id": "75a57d03620fe9d07c9b2f275b77e745f22b7812",
          "url": "https://github.com/CGI-FR/posimap/commit/9f283e822e2cd0c26dce02b7ccbe504ec6dbd3ad"
        },
        "date": 1747063273915,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 1395075,
            "unit": "ns/op\t 1067839 B/op\t    6620 allocs/op",
            "extra": "8515 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 1395075,
            "unit": "ns/op",
            "extra": "8515 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 1067839,
            "unit": "B/op",
            "extra": "8515 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 6620,
            "unit": "allocs/op",
            "extra": "8515 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold",
            "value": 895397,
            "unit": "ns/op\t  682641 B/op\t    3703 allocs/op",
            "extra": "13468 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - ns/op",
            "value": 895397,
            "unit": "ns/op",
            "extra": "13468 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - B/op",
            "value": 682641,
            "unit": "B/op",
            "extra": "13468 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - allocs/op",
            "value": 3703,
            "unit": "allocs/op",
            "extra": "13468 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "545113b51edb7c442e2b3767f09a349b6e2a9f33",
          "message": "fix: buffer too short should fail (#32)\n\n* fix: eof with short line does not clean buffer\n\n* style: lint",
          "timestamp": "2025-05-14T11:31:35+02:00",
          "tree_id": "574f74110024b7eb5f140ddec668c4011da59b8e",
          "url": "https://github.com/CGI-FR/posimap/commit/545113b51edb7c442e2b3767f09a349b6e2a9f33"
        },
        "date": 1747215209765,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 1368502,
            "unit": "ns/op\t 1067825 B/op\t    6620 allocs/op",
            "extra": "8497 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 1368502,
            "unit": "ns/op",
            "extra": "8497 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 1067825,
            "unit": "B/op",
            "extra": "8497 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 6620,
            "unit": "allocs/op",
            "extra": "8497 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold",
            "value": 874643,
            "unit": "ns/op\t  682641 B/op\t    3703 allocs/op",
            "extra": "13714 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - ns/op",
            "value": 874643,
            "unit": "ns/op",
            "extra": "13714 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - B/op",
            "value": 682641,
            "unit": "B/op",
            "extra": "13714 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - allocs/op",
            "value": 3703,
            "unit": "allocs/op",
            "extra": "13714 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "adrien.aury@cgi.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "adrien.aury@cgi.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "distinct": true,
          "id": "16fe22585fb188390fe6cdbcda06077c34a09d29",
          "message": "refactor: add missing methods in buffer and document api",
          "timestamp": "2025-05-14T09:48:15Z",
          "tree_id": "df6c6da03f5f70f0afa298abb4d524dbd0ce2cb6",
          "url": "https://github.com/CGI-FR/posimap/commit/16fe22585fb188390fe6cdbcda06077c34a09d29"
        },
        "date": 1747216222789,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 1370211,
            "unit": "ns/op\t 1067842 B/op\t    6620 allocs/op",
            "extra": "8457 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 1370211,
            "unit": "ns/op",
            "extra": "8457 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 1067842,
            "unit": "B/op",
            "extra": "8457 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 6620,
            "unit": "allocs/op",
            "extra": "8457 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold",
            "value": 883790,
            "unit": "ns/op\t  682644 B/op\t    3703 allocs/op",
            "extra": "13569 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - ns/op",
            "value": 883790,
            "unit": "ns/op",
            "extra": "13569 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - B/op",
            "value": 682644,
            "unit": "B/op",
            "extra": "13569 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - allocs/op",
            "value": 3703,
            "unit": "allocs/op",
            "extra": "13569 times\n4 procs"
          }
        ]
      },
      {
        "commit": {
          "author": {
            "email": "44274230+adrienaury@users.noreply.github.com",
            "name": "Adrien Aury",
            "username": "adrienaury"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "4e61ef1cb5f210bb4ac829a3e0bfeb80fb4b31e5",
          "message": "feat: configure records separator (#33)\n\n* feat: records separator\n\n* feat: records separator\n\n* feat: records separator\n\n* feat: records separator\n\n* feat: require length or separator to be set",
          "timestamp": "2025-05-14T15:08:09+02:00",
          "tree_id": "d0cafb2a6343ca8c1727a6c42a6781ad15eb30c6",
          "url": "https://github.com/CGI-FR/posimap/commit/4e61ef1cb5f210bb4ac829a3e0bfeb80fb4b31e5"
        },
        "date": 1747228204184,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkFold",
            "value": 1365133,
            "unit": "ns/op\t 1068631 B/op\t    6640 allocs/op",
            "extra": "8508 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - ns/op",
            "value": 1365133,
            "unit": "ns/op",
            "extra": "8508 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - B/op",
            "value": 1068631,
            "unit": "B/op",
            "extra": "8508 times\n4 procs"
          },
          {
            "name": "BenchmarkFold - allocs/op",
            "value": 6640,
            "unit": "allocs/op",
            "extra": "8508 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold",
            "value": 883637,
            "unit": "ns/op\t  683565 B/op\t    3725 allocs/op",
            "extra": "13622 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - ns/op",
            "value": 883637,
            "unit": "ns/op",
            "extra": "13622 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - B/op",
            "value": 683565,
            "unit": "B/op",
            "extra": "13622 times\n4 procs"
          },
          {
            "name": "BenchmarkUnfold - allocs/op",
            "value": 3725,
            "unit": "allocs/op",
            "extra": "13622 times\n4 procs"
          }
        ]
      }
    ]
  }
}