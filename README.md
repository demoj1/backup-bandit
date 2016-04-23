## Пример файла конфигурации config.yaml
---
```yaml
email_settings:
  login: secret_email@yandex.ru
  password: "*********"
email_list:
  - dmitryd.prog@gmail.com
  - example2@mail.ru
  - ...
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

## Файл конфигурации для сбора логов robber.yaml
---
```yaml
tools:
  -
    path: /usr/bin/discus
    args: -c
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
  -
    path: /usr/bin/uptime
    groups:
      - up
      - users
      - avr
    regex: >-
        up
        (?P<up>.*),\s*
        (?P<users>\d+) users,\s*load average:
        (?P<avr>.*\d)
```
## Сообщение на почту
![Сообщение](http://storage1.static.itmages.ru/i/16/0423/h_1461432682_2094362_119a6cf381.png)
