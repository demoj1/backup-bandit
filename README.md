## Пример файла конфигурации
---
```yaml
paths:
  -
    path: "%[y]/%[m]/%[-1d]/secret_path"
    files:
      -
        name: "secret_path2[.]bak"
        min_valid_size: 30 MB
      -
        name: "secret_path3[.]bak"
        min_valid_size: 200 MB
  -
    path: "%[y]/%[m]/%[d]/secret_path4"
    files:
      -
        name: ".*_secret_path4[.]tar"
        min_valid_size: 600 MB
  -
    path: "%[y]/%[m]/%[d]/secret_path5/docs"
    min_valid_size: 600 MB
```
Временные метки строчными буквами обозначают формат без ведущего нуля.

Например строка: `"%[y]/%[m]/%[d]/secret_path5/docs"` будет иметь эквивалент: `"2016/4/1/secret_path5/docs"`

Прописными буквами: `"%[Y]/%[M]/%[D]/secret_path5/docs"` эквивалента: `"2016/04/01/secret_path5/docs"`

Минимальный размер файла может содержать разные обозначения:
  * KB
  * MB
  * GB

## Файл конфигурации для сбора логов
---
```yaml
tools:
  -
    path: /usr/bin/discus
    groups:
      - mount
      - total
      - used
      - avail
      - prcnt
      - graph
    regex: >-
        (?P<mount>\/.*?)\s+
        (?P<total>\d+[.]?\d* \w+[B])\s+
        (?P<used>\d+[.]?\d* \w+[B])\s+
        (?P<avail>\d+[.]?\d* \w+[B])\s+
        (?P<prcnt>\d+[.]?\d*%)\s+
        (?P<graph>\[.*\])
```
Вывод discus
```
Mount           Total         Used         Avail      Prcnt      Graph
/run            767.7 MB      76.9 MB     690.8 MB    10.0%   [*---------]
/              109.17 GB      8.75 GB    100.43 GB     8.0%   [*---------]
+l/security         0 KB         0 KB         0 KB     0.0%   [----------]
/dev/shm         3.75 GB      98.8 MB      3.65 GB     2.6%   [----------]
/run/lock         5.0 MB         4 KB       5.0 MB     0.1%   [----------]
+/fs/cgroup      3.75 GB         0 KB      3.75 GB     0.0%   [----------]

```
