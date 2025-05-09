window.BENCHMARK_DATA = {
  "lastUpdate": 1746792138308,
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
      }
    ]
  }
}