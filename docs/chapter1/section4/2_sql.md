# SQLで遊ぶ

:::tip
SQL 文は小文字でも動きます。大文字を打つのが面倒な場合は、小文字にしましょう
:::

データベースに入った状態で、さまざまな SQL を試していきます。ターミナルに `mysql>`が表示されていることを確認してください。

初めにデータベースを確認します。

```sql
mysql> SHOW DATABASES;
```

以下のように出力されるはずです。

```txt
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
| world              |
+--------------------+
5 rows in set (0.00 sec)
```

これは MySQL に含まれているデータベースの一覧です。今回は`world`というデータベースを使います。下のコマンドを入力してください。

```sql
mysql> USE world;
```

他のデータベースは MySQL の設定などが含まれているデータベースです。
`world`というデータベースは、 MySQL 公式が用意している学習用のデータベースで、さまざまな国や地域の情報が集まっています。

https://dev.mysql.com/doc/world-setup/en/

次に`world`に含まれるテーブル一覧を確認します。

```sql
mysql> SHOW TABLES;
```

```txt
+-----------------+
| Tables_in_world |
+-----------------+
| city            |
| country         |
| countrylanguage |
+-----------------+
3 rows in set (0.00 sec)
```

`city`、`country`、`countrylanguage`の 3 つのテーブルの存在がわかります。

## テーブルの構造を見る

```sql
mysql> DESC city;
-- または
mysql> DESCRIBE city;
-- または
mysql> SHOW COLUMNS FROM city;
```

```txt
+-------------+----------+------+-----+---------+----------------+
| Field       | Type     | Null | Key | Default | Extra          |
+-------------+----------+------+-----+---------+----------------+
| ID          | int(11)  | NO   | PRI | NULL    | auto_increment |
| Name        | char(35) | NO   |     |         |                |
| CountryCode | char(3)  | NO   | MUL |         |                |
| District    | char(20) | NO   |     |         |                |
| Population  | int(11)  | NO   |     | 0       |                |
+-------------+----------+------+-----+---------+----------------+
5 rows in set (0.00 sec)
```

`city`テーブルには`ID`、`Name`、`CountryCode`、`District`、`Population`の 5 つのカラムの存在がわかります。

## データベースから情報を取得する

### SELECT 文

`SELECT 対象カラム名 FROM 対象テーブル名;` で、テーブルから情報を取得できます。複数のカラムを取得したいときは`,`で区切ります。

```sql
mysql> SELECT Name, Population FROM city;
```

```txt
+-----------------------------------+------------+
| Name                              | Population |
+-----------------------------------+------------+
| Kabul                             |    1780000 |
| Qandahar                          |     237500 |
| Herat                             |     186800 |
| Mazar-e-Sharif                    |     127800 |
| Amsterdam                         |     731200 |
| Rotterdam                         |     593321 |
| Haag                              |     440900 |
| Utrecht                           |     234323 |
| Eindhoven                         |     201843 |
| Tilburg                           |     193238 |
...省略
```

#### 全取得

全てのカラムを取得したい場合は、`*`を使えます。

```sql
mysql> SELECT * FROM city;
```

```txt
+------+-----------------------------------+-------------+----------------------+------------+
| ID   | Name                              | CountryCode | District             | Population |
+------+-----------------------------------+-------------+----------------------+------------+
|    1 | Kabul                             | AFG         | Kabol                |    1780000 |
|    2 | Qandahar                          | AFG         | Qandahar             |     237500 |
|    3 | Herat                             | AFG         | Herat                |     186800 |
|    4 | Mazar-e-Sharif                    | AFG         | Balkh                |     127800 |
|    5 | Amsterdam                         | NLD         | Noord-Holland        |     731200 |
|    6 | Rotterdam                         | NLD         | Zuid-Holland         |     593321 |
|    7 | Haag                              | NLD         | Zuid-Holland         |     440900 |
|    8 | Utrecht                           | NLD         | Utrecht              |     234323 |
|    9 | Eindhoven                         | NLD         | Noord-Brabant        |     201843 |
|   10 | Tilburg                           | NLD         | Noord-Brabant        |     193238 |
...省略
```

#### LIMIT句

SELECT 文の後ろに`LIMIT 件数`を追加することで取得件数の上限を指定できます。

```sql
mysql> SELECT * FROM city LIMIT 5;
```

```txt
+----+----------------+-------------+---------------+------------+
| ID | Name           | CountryCode | District      | Population |
+----+----------------+-------------+---------------+------------+
|  1 | Kabul          | AFG         | Kabol         |    1780000 |
|  2 | Qandahar       | AFG         | Qandahar      |     237500 |
|  3 | Herat          | AFG         | Herat         |     186800 |
|  4 | Mazar-e-Sharif | AFG         | Balkh         |     127800 |
|  5 | Amsterdam      | NLD         | Noord-Holland |     731200 |
+----+----------------+-------------+---------------+------------+
5 rows in set (0.00 sec)
```

#### OFFSET句

`OFFSET ずらす数`を LIMIT 句の後ろにつなげると、データを取得し始める位置をずらして指定できます。`LIMIT`を指定せずに`OFFSET`を指定できません。

```sql
mysql> SELECT * FROM city LIMIT 5 OFFSET 10;
```

```txt
+----+-----------+-------------+---------------+------------+
| ID | Name      | CountryCode | District      | Population |
+----+-----------+-------------+---------------+------------+
| 11 | Groningen | NLD         | Groningen     |     172701 |
| 12 | Breda     | NLD         | Noord-Brabant |     160398 |
| 13 | Apeldoorn | NLD         | Gelderland    |     153491 |
| 14 | Nijmegen  | NLD         | Gelderland    |     152463 |
| 15 | Enschede  | NLD         | Overijssel    |     149544 |
+----+-----------+-------------+---------------+------------+
5 rows in set (0.00 sec)
```

#### WHERE句

`SELECT カラム名 FROM テーブル名 WHERE 条件式;`で取得するレコードの条件を付けることができます。`AND`や`OR`を使うと条件を複数つけることができます。

```sql
mysql> SELECT * FROM city WHERE Population >= 8000000;
```

```txt
+------+------------------+-------------+------------------+------------+
| ID   | Name             | CountryCode | District         | Population |
+------+------------------+-------------+------------------+------------+
|  206 | S�o Paulo        | BRA         | S�o Paulo        |    9968485 |
|  939 | Jakarta          | IDN         | Jakarta Raya     |    9604900 |
| 1024 | Mumbai (Bombay)  | IND         | Maharashtra      |   10500000 |
| 1890 | Shanghai         | CHN         | Shanghai         |    9696300 |
| 2331 | Seoul            | KOR         | Seoul            |    9981619 |
| 2515 | Ciudad de M�xico | MEX         | Distrito Federal |    8591309 |
| 2822 | Karachi          | PAK         | Sindh            |    9269265 |
| 3357 | Istanbul         | TUR         | Istanbul         |    8787958 |
| 3580 | Moscow           | RUS         | Moscow (City)    |    8389200 |
| 3793 | New York         | USA         | New York         |    8008278 |
+------+------------------+-------------+------------------+------------+
10 rows in set (0.01 sec)
```

```sql
mysql> SELECT * FROM city WHERE CountryCode = "JPN" AND Population > 5000000;
```

```txt
+------+-------+-------------+----------+------------+
| ID   | Name  | CountryCode | District | Population |
+------+-------+-------------+----------+------------+
| 1532 | Tokyo | JPN         | Tokyo-to |    7980230 |
+------+-------+-------------+----------+------------+
1 row in set (0.00 sec)
```

#### ORDER BY 句

`SELECT カラム名 FROM テーブル名 ORDERED BY 対象カラム名 並び順;`で結果を昇順・降順に並び替えて取得できます。`ASC`で昇順、`DESC`で降順です。

```sql
SELECT * FROM city WHERE Population >= 8000000 ORDER BY Population DESC;
```

```txt
+------+------------------+-------------+------------------+------------+
| ID   | Name             | CountryCode | District         | Population |
+------+------------------+-------------+------------------+------------+
| 1024 | Mumbai (Bombay)  | IND         | Maharashtra      |   10500000 |
| 2331 | Seoul            | KOR         | Seoul            |    9981619 |
|  206 | S�o Paulo        | BRA         | S�o Paulo        |    9968485 |
| 1890 | Shanghai         | CHN         | Shanghai         |    9696300 |
|  939 | Jakarta          | IDN         | Jakarta Raya     |    9604900 |
| 2822 | Karachi          | PAK         | Sindh            |    9269265 |
| 3357 | Istanbul         | TUR         | Istanbul         |    8787958 |
| 2515 | Ciudad de M�xico | MEX         | Distrito Federal |    8591309 |
| 3580 | Moscow           | RUS         | Moscow (City)    |    8389200 |
| 3793 | New York         | USA         | New York         |    8008278 |
+------+------------------+-------------+------------------+------------+
10 rows in set (0.01 sec)
```

#### IN演算子

`SELECT カラム名 FROM テーブル名 WHERE カラム名 IN (データ1, データ2, ...)`のように書くことで、カラムの値が複数の値のうちどれかに当てはまるものを選ぶことができます。
例えば、都市のうち都道府県(`District`)が四国(香川、徳島、愛媛、高知)に当てはまるものを選ぶ文は下のようになります。

```sql
mysql> SELECT * FROM city WHERE District IN ("Kagawa", "Tokushima", "Ehime", "Kochi");
```

```txt
+------+-----------+-------------+-----------+------------+
| ID   | Name      | CountryCode | District  | Population |
+------+-----------+-------------+-----------+------------+
| 1559 | Matsuyama | JPN         | Ehime     |     466133 |
| 1586 | Takamatsu | JPN         | Kagawa    |     332471 |
| 1592 | Kochi     | JPN         | Kochi     |     324710 |
| 1610 | Tokushima | JPN         | Tokushima |     269649 |
| 1700 | Niihama   | JPN         | Ehime     |     127207 |
| 1716 | Imabari   | JPN         | Ehime     |     119357 |
+------+-----------+-------------+-----------+------------+
6 rows in set (0.01 sec)
```

#### JOIN句

`SELECT カラム名 FROM テーブル名1 JOIN テーブル名2 ON 条件式;`で複数のテーブルを結合して、1 つのテーブルとして取得できます。
条件式は、結合したいテーブルの特定のカラムの関係について書きます。

中国語を使っている国の国名を知りたいときを考えましょう。
`countrylanguage`テーブルには下のように国コード(`CountryCode`)のカラムはありますが国名はありません。そのため、`country`テーブルから国名を知る必要があります。

```sql
mysql> DESC countrylanguage;
```

```txt
+-------------+---------------+------+-----+---------+-------+
| Field       | Type          | Null | Key | Default | Extra |
+-------------+---------------+------+-----+---------+-------+
| CountryCode | char(3)       | NO   | PRI |         |       |
| Language    | char(30)      | NO   | PRI |         |       |
| IsOfficial  | enum('T','F') | NO   |     | F       |       |
| Percentage  | float(4,1)    | NO   |     | 0.0     |       |
+-------------+---------------+------+-----+---------+-------+
4 rows in set (0.00 sec)
```

そこで`JOIN`句を下のように使います。

```sql
mysql> SELECT country.Name, countrylanguage.Language FROM country JOIN countrylanguage ON country.Code = countrylanguage.CountryCode WHERE countrylanguage.Language = "Chinese";
```

```txt
+--------------------------+----------+
| Name                     | Language |
+--------------------------+----------+
| Brunei                   | Chinese  |
| Canada                   | Chinese  |
| China                    | Chinese  |
| Costa Rica               | Chinese  |
| Christmas Island         | Chinese  |
| Japan                    | Chinese  |
| Cambodia                 | Chinese  |
| South Korea              | Chinese  |
| Northern Mariana Islands | Chinese  |
| Malaysia                 | Chinese  |
| Nauru                    | Chinese  |
| Palau                    | Chinese  |
| North Korea              | Chinese  |
| French Polynesia         | Chinese  |
| R�union                  | Chinese  |
| Singapore                | Chinese  |
| Thailand                 | Chinese  |
| United States            | Chinese  |
| Vietnam                  | Chinese  |
+--------------------------+----------+
19 rows in set (0.00 sec)
```

`country`テーブルの`Code`カラムの値と`countrylanguage`テーブルの`CountryCode`カラムが一致するように 2 つのテーブルをつなげています。`WHERE`句は`JOIN`句の後に書く必要があることに注意しましょう。

JOIN 句にはいくつか種類があり、適切なものを使う必要があります。今回使ったものは`INNER JOIN`と呼ばれます。

https://www.w3schools.com/sql/sql_join.asp

#### AS句

`SELECT カラム名 AS 別名 FROM テーブル名`で、カラムに別名を付けて扱うことができます。
例えば日本の都市の名前(`Name`)と都道府県(`District`)を取得したいとき、`District`を`Prefecture`と表示したい場合は次のように書くことができます。

```sql
mysql> SELECT Name, District AS "Prefecture" FROM city WHERE CountryCode = "JPN";
```

```txt
+---------------------+------------+
| Name                | Prefecture |
+---------------------+------------+
| Tokyo               | Tokyo-to   |
| Jokohama [Yokohama] | Kanagawa   |
| Osaka               | Osaka      |
| Nagoya              | Aichi      |
| Sapporo             | Hokkaido   |
| Kioto               | Kyoto      |
| Kobe                | Hyogo      |
| Fukuoka             | Fukuoka    |
| Kawasaki            | Kanagawa   |
| Hiroshima           | Hiroshima  |
...省略
```

`AS`はカラム名だけでなくテーブル名にも使うことができ、先ほどの`JOIN`の SQL は`AS`を使うとこのように書けます。

```sql
mysql> SELECT c.Name, cl.Language FROM country AS "c" JOIN countrylanguage AS "cl" ON c.Code = cl.CountryCode WHERE cl.Language = "Chinese";
```

また、`AS`は省略できます。

```sql
mysql> SELECT Name, District "Prefecture" FROM city WHERE CountryCode = "JPN";
```

#### COUNT関数

`SELECT COUNT(カラム名) FROM テーブル名;`でレコードの数を数えることができます。
都市のうち国コード(`CountryCode`)が`JPN`のレコード数は下のようにして取得できます。

```sql
mysql> SELECT COUNT(*) FROM city WHERE CountyCode = "JPN";
```

```txt
+----------+
| count(*) |
+----------+
|      248 |
+----------+
1 row in set (0.00 sec)
```

#### GROUP BY 句

`GROUP BY カラム名`を付けることで、`COUNT`などの結果を共通の値でまとめることができます。
各国コードの都市数を数える SQL 文は下のようになります。

```sql
SELECT CountryCode, COUNT(*) FROM city GROUP BY CountryCode;
```

```txt
+-------------+----------+
| CountryCode | COUNT(*) |
+-------------+----------+
| ABW         |        1 |
| AFG         |        4 |
| AGO         |        5 |
| AIA         |        2 |
| ALB         |        1 |
| AND         |        1 |
| ANT         |        1 |
| ARE         |        5 |
| ARG         |       57 |
...省略
```

## データベースの値を変える

### INSERT文

`INSERT INTO テーブル名 (カラム名1, カラム名2, ...) VALUES (値1, 値2, ...);`でテーブルにレコードを挿入できます。

1. 挿入

```sql
mysql> INSERT INTO city (Name, CountryCode, District, Population) VALUES ("oookayama", "JPN", "Tokyo-to", 5000);
```

2. 確認

```sql
mysql> SELECT * FROM city ORDER BY ID DESC LIMIT 1;
```

```txt
+------+-----------+-------------+----------+------------+
| ID   | Name      | CountryCode | District | Population |
+------+-----------+-------------+----------+------------+
| 4080 | oookayama | JPN         | Tokyo-to |       5000 |
+------+-----------+-------------+----------+------------+
1 row in set (0.00 sec)
```

### UPDATE文

`UPDATE テーブル名 SET カラム名 = 値 WHERE 条件式;`で条件に当てはまるレコードの値を変えることができます。
さっき追加した大岡山の情報を変えてみましょう。

```sql
mysql> UPDATE city SET Population = 9999 WHERE ID = 4080;
mysql> SELECT * FROM city WHERE ID = 4080;
```

```txt
+------+-----------+-------------+----------+------------+
| ID   | Name      | CountryCode | District | Population |
+------+-----------+-------------+----------+------------+
| 4080 | oookayama | JPN         | Tokyo-to |       9999 |
+------+-----------+-------------+----------+------------+
1 row in set (0.00 sec)
```

### DELETE文

`DELETE FROM テーブル名 WHERE 条件式;`の構文で条件に合致するレコードを削除できます。
大岡山を消してみましょう。

```sql
mysql> DELETE FROM city WHERE ID = 4080;
```

確認してみます。

```sql
mysql> SELECT * FROM city WHERE ID = 4080;
```

```txt
Empty set (0.00 sec)
```

大岡山が消えていることが確認できました。

## Adminerを使う

Adminer(https://www.adminer.org/) はデータベースを GUI（マウスなど）を使って操作するためのソフトウェアです。traP 内では traQ の開発などで使われています。同じようなソフトウェアとして PHPMyAdmin などがあります。これらを使うことで SQL を使わなくてもデータベースを操作できます。

今回は`task up`を実行したときに Adminer が立ち上がるようになっています。 http://localhost:8080 にアクセスすると使えます。
ログイン画面が出てくるはずなので、MySQL にログインするときと同様に、下の画像のように入力してログインしてください。パスワードは`password`です。

![](assets/adminer_login.png)

![](assets/adminer_home.png)

ログインした画面からテーブルを選び、「データ」を選択するとレコード一覧を見ることができ、検索や並び替えなどができます。また、データの編集もできます。Adminer 上で SQL の実行もでき、文がいい感じに強調表示されるので書きやすいです。

![](assets/adminer_sql.png)

## 演習問題

基本編・応用編があります。基本編はここまで説明したことだけで解けるはずです。応用編はここまで紹介していないものも使うので、自分で調べて解く必要があります。基本編は解けるようにしましょう。応用編は余裕があったらやってみてください。
SQL を実行するのは Adminer 上でも、`task db`を実行して MySQL にログインしてもどちらでも自分が好きな方でやってください。

### 基本編

#### 1-1

`country`テーブルから、日本(`Name`のカラムが`Japan`)の情報を全て取得してください。

:::details **答え**

```sql
SELECT * FROM country WHERE Name = "Japan";
```

**出力**

```txt
+------+-------+-----------+--------------+-------------+-----------+------------+----------------+------------+------------+--------------+-------------------------+-------------+---------+-------+
| Code | Name  | Continent | Region       | SurfaceArea | IndepYear | Population | LifeExpectancy | GNP        | GNPOld     | LocalName    | GovernmentForm          | HeadOfState | Capital | Code2 |
+------+-------+-----------+--------------+-------------+-----------+------------+----------------+------------+------------+--------------+-------------------------+-------------+---------+-------+
| JPN  | Japan | Asia      | Eastern Asia |   377829.00 |      -660 |  126714000 |           80.7 | 3787042.00 | 4192638.00 | Nihon/Nippon | Constitutional Monarchy | Akihito     |    1532 | JP    |
+------+-------+-----------+--------------+-------------+-----------+------------+----------------+------------+------------+--------------+-------------------------+-------------+---------+-------+
1 row in set (0.04 sec)
```

:::

#### 1-2

`country`テーブルから、独立年（建国年）(`IndepYear`)が`0`以下の国の情報をすべて取得してください。

:::details **答え**

```sql
SELECT * FROM country WHERE IndepYear <= 0;
```

**出力**

```txt
+------+----------+-----------+----------------+-------------+-----------+------------+----------------+------------+------------+--------------+-------------------------+----------------+---------+-------+
| Code | Name     | Continent | Region         | SurfaceArea | IndepYear | Population | LifeExpectancy | GNP        | GNPOld     | LocalName    | GovernmentForm          | HeadOfState    | Capital | Code2 |
+------+----------+-----------+----------------+-------------+-----------+------------+----------------+------------+------------+--------------+-------------------------+----------------+---------+-------+
| CHN  | China    | Asia      | Eastern Asia   |  9572900.00 |     -1523 | 1277558000 |           71.4 |  982268.00 |  917719.00 | Zhongquo     | People'sRepublic        | Jiang Zemin    |    1891 | CN    |
| ETH  | Ethiopia | Africa    | Eastern Africa |  1104300.00 |     -1000 |   62565000 |           45.2 |    6353.00 |    6180.00 | YeItyop�iya  | Republic                | Negasso Gidada |     756 | ET    |
| JPN  | Japan    | Asia      | Eastern Asia   |   377829.00 |      -660 |  126714000 |           80.7 | 3787042.00 | 4192638.00 | Nihon/Nippon | Constitutional Monarchy | Akihito        |    1532 | JP    |
+------+----------+-----------+----------------+-------------+-----------+------------+----------------+------------+------------+--------------+-------------------------+----------------+---------+-------+
3 rows in set (0.00 sec)
```

:::

#### 1-3

`country`テーブルから、大陸(`Continent`)がアジア(`Asia`)の国のうち、人口(`Population`)の上位 5 か国の国名(`Name`)と人口を多い順に取得してください。

:::details **答え**

```sql
SELECT Name, Population FROM country WHERE Continent = "Asia" ORDER BY Population DESC LIMIT 5;
```

**出力**

```txt
+------------+------------+
| Name       | Population |
+------------+------------+
| China      | 1277558000 |
| India      | 1013662000 |
| Indonesia  |  212107000 |
| Pakistan   |  156483000 |
| Bangladesh |  129155000 |
+------------+------------+
5 rows in set (0.00 sec)
```

:::

#### 1-4

`country`テーブルから人口が多い順に 11 位から 15 位の国の情報をすべて取得してください。

:::details **答え**

```sql
SELECT * FROM country ORDER BY Population DESC LIMIT 5 OFFSET 10;
```

**出力**

```txt
+------+-------------+---------------+-----------------+-------------+-----------+------------+----------------+------------+------------+-------------+----------------------+-------------------------+---------+-------+
| Code | Name        | Continent     | Region          | SurfaceArea | IndepYear | Population | LifeExpectancy | GNP        | GNPOld     | LocalName   | GovernmentForm       | HeadOfState             | Capital | Code2 |
+------+-------------+---------------+-----------------+-------------+-----------+------------+----------------+------------+------------+-------------+----------------------+-------------------------+---------+-------+
| MEX  | Mexico      | North America | Central America |  1958201.00 |      1810 |   98881000 |           71.5 |  414972.00 |  401461.00 | M�xico      | Federal Republic     | Vicente Fox Quesada     |    2515 | MX    |
| DEU  | Germany     | Europe        | Western Europe  |   357022.00 |      1955 |   82164700 |           77.4 | 2133367.00 | 2102826.00 | Deutschland | Federal Republic     | Johannes Rau            |    3068 | DE    |
| VNM  | Vietnam     | Asia          | Southeast Asia  |   331689.00 |      1945 |   79832000 |           69.3 |   21929.00 |   22834.00 | Vi�t Nam    | Socialistic Republic | Tr�n Duc Luong          |    3770 | VN    |
| PHL  | Philippines | Asia          | Southeast Asia  |   300000.00 |      1946 |   75967000 |           67.5 |   65107.00 |   82239.00 | Pilipinas   | Republic             | Gloria Macapagal-Arroyo |     766 | PH    |
| EGY  | Egypt       | Africa        | Northern Africa |  1001449.00 |      1922 |   68470000 |           63.3 |   82710.00 |   75617.00 | Misr        | Republic             | Hosni Mubarak           |     608 | EG    |
+------+-------------+---------------+-----------------+-------------+-----------+------------+----------------+------------+------------+-------------+----------------------+-------------------------+---------+-------+
5 rows in set (0.00 sec)
```

:::

#### 1-5

`countrylanguage`テーブルから、言語(`Language`)が日本語(`Japanese`)である国を数えて、数を`Japanese`という項目名で取得してください。

:::details **答え**

```sql
SELECT COUNT(*) AS "Japanese" FROM countrylanguage WHERE Language = "Japanese";
```

**出力**

```txt
+----------+
| Japanese |
+----------+
|        4 |
+----------+
1 row in set (0.04 sec)
```

:::

#### 1-6

`country`テーブルから、大陸(`Continent`)ごとの国の数を数えて大陸ごとに取得してください。

:::details **答え**

```sql
SELECT Continent, COUNT(*) FROM country GROUP BY Continent;
```

**出力**

```txt
+---------------+----------+
| Continent     | COUNT(*) |
+---------------+----------+
| North America |       37 |
| Asia          |       51 |
| Africa        |       58 |
| Europe        |       46 |
| South America |       14 |
| Oceania       |       28 |
| Antarctica    |        5 |
+---------------+----------+
7 rows in set (0.01 sec)
```

:::

#### 1-7

`city`テーブルと`country`テーブルを使って、国名と首都の一覧を、国名を`Country`、首都を`Capital`として表示してください。

ヒント：`country`テーブルの`Capital`カラムは、その国の首都の`city`テーブルでの`ID`を指しています。

:::details **答え**

```sql
SELECT country.Name AS "Country", city.Name AS "Capital" FROM country JOIN city ON country.Capital = city.ID;
```

**出力**

```txt
+---------------------------------------+-----------------------------------+
| Country                               | Capital                           |
+---------------------------------------+-----------------------------------+
| Aruba                                 | Oranjestad                        |
| Afghanistan                           | Kabul                             |
| Angola                                | Luanda                            |
| Anguilla                              | The Valley                        |
| Albania                               | Tirana                            |
| Andorra                               | Andorra la Vella                  |
| Netherlands Antilles                  | Willemstad                        |
| United Arab Emirates                  | Abu Dhabi                         |
| Argentina                             | Buenos Aires                      |
| Armenia                               | Yerevan                           |
| American Samoa                        | Fagatogo                          |
| Antigua and Barbuda                   | Saint John�s                      |
| Australia                             | Canberra                          |
| Austria                               | Wien                              |
| Azerbaijan                            | Baku                              |
| Burundi                               | Bujumbura                         |
省略...
```

:::

#### 1-8

`country`テーブルから、大陸(`Continent`)が北アメリカ(`North America`)または南アメリカ(`South America`)の国の情報をすべて取得してください。この問題は答えが 2 通りあります。余裕のある人は 2 つ考えてみましょう。

:::details **答え**

```sql
SELECT * FROM country WHERE Continent = "North America" OR Continent = "South America";
--または
SELECT * FROM country WHERE Continent IN ("North America", "South America");
```

**出力**

```txt
+------+----------------------+---------------+---------------+-------------+-----------+------------+----------------+-----------+-----------+----------------------+----------------------------------------------+--------------------+---------+-------+
| Code | Name                 | Continent     | Region        | SurfaceArea | IndepYear | Population | LifeExpectancy | GNP       | GNPOld    | LocalName            | GovernmentForm                               | HeadOfState        | Capital | Code2 |
+------+----------------------+---------------+---------------+-------------+-----------+------------+----------------+-----------+-----------+----------------------+----------------------------------------------+--------------------+---------+-------+
| ABW  | Aruba                | North America | Caribbean     |      193.00 |      NULL |     103000 |           78.4 |    828.00 |    793.00 | Aruba                | Nonmetropolitan Territory of The Netherlands | Beatrix            |     129 | AW    |
| AIA  | Anguilla             | North America | Caribbean     |       96.00 |      NULL |       8000 |           76.1 |     63.20 |      NULL | Anguilla             | Dependent Territory of the UK                | Elisabeth II       |      62 | AI    |
| ANT  | Netherlands Antilles | North America | Caribbean     |      800.00 |      NULL |     217000 |           74.7 |   1941.00 |      NULL | Nederlandse Antillen | Nonmetropolitan Territory of The Netherlands | Beatrix            |      33 | AN    |
| ARG  | Argentina            | South America | South America |  2780400.00 |      1816 |   37032000 |           75.1 | 340238.00 | 323310.00 | Argentina            | Federal Republic                             | Fernando de la R�a |      69 | AR    |
| ATG  | Antigua and Barbuda  | North America | Caribbean     |      442.00 |      1981 |      68000 |           70.5 |    612.00 |    584.00 | Antigua and Barbuda  | Constitutional Monarchy                      | Elisabeth II       |      63 | AG    |
...省略
```

:::

#### 1-9

`country`テーブルに以下の情報を持った国を挿入してください。

- 国名 `traP`
- 国コード `TRP`
- 大陸 `Asia`
- 地域 `Eastern Asia`
- 現地での呼び名(`LocalName`) `traP`
- 政治体制(`GovernmentForm`) 立憲君主制(`Constitutional Monarchy`)
- 人口 `500`人
- 2 文字国コード(`Code2`) `TP`

:::details **答え**

```sql
INSERT INTO country (Name, Code, Continent, Region, LocalName, GovernmentForm, Population, Code2) VALUES ("traP", "TRP", "Asia", "Eastern Asia", "traP", "Constitutional Monarchy", 500, "TP");
```

**確認用**

```sql
SELECT * FROM country WHERE Name = "traP";
```

**出力**

```txt
+------+------+-----------+--------------+-------------+-----------+------------+----------------+------+--------+-----------+-------------------------+-------------+---------+-------+
| Code | Name | Continent | Region       | SurfaceArea | IndepYear | Population | LifeExpectancy | GNP  | GNPOld | LocalName | GovernmentForm          | HeadOfState | Capital | Code2 |
+------+------+-----------+--------------+-------------+-----------+------------+----------------+------+--------+-----------+-------------------------+-------------+---------+-------+
| TRP  | traP | Asia      | Eastern Asia |        0.00 |      NULL |        500 |           NULL | NULL |   NULL | traP      | Constitutional Monarchy | NULL        |    NULL | TP    |
+------+------+-----------+--------------+-------------+-----------+------------+----------------+------+--------+-----------+-------------------------+-------------+---------+-------+
1 row in set (0.00 sec)
```

:::

#### 1-10

1-9 で作った国「`traP`」の独立年(`IndepYear`)を 2015 年にしましょう。

:::details **答え**

```sql
UPDATE country SET IndepYear = 2015 WHERE Code = "TRP";
```

`Code`で指定しなくても、国を 1 つに指定できれば他のカラムで条件式を作ってもよいです(`Name="traP"`など)。

**確認用**

```sql
SELECT * FROM country WHERE Code = "TRP";
```

**出力**

```txt
+------+------+-----------+--------------+-------------+-----------+------------+----------------+------+--------+-----------+-------------------------+-------------+---------+-------+
| Code | Name | Continent | Region       | SurfaceArea | IndepYear | Population | LifeExpectancy | GNP  | GNPOld | LocalName | GovernmentForm          | HeadOfState | Capital | Code2 |
+------+------+-----------+--------------+-------------+-----------+------------+----------------+------+--------+-----------+-------------------------+-------------+---------+-------+
| TRP  | traP | Asia      | Eastern Asia |        0.00 |      2015 |        500 |           NULL | NULL |   NULL | traP      | Constitutional Monarchy | NULL        |    NULL | TP    |
+------+------+-----------+--------------+-------------+-----------+------------+----------------+------+--------+-----------+-------------------------+-------------+---------+-------+
1 row in set (0.00 sec)
```

:::

#### 1-11

`country`に追加した国「`traP`」を削除しましょう。

:::details **答え**

```sql
DELETE FROM country WHERE Code = "TRP";
```

**確認用**

```sql
SELECT * FROM country WHERE Code = "TRP";
```

**出力**

```txt
Empty set (0.00 sec)
```

:::

### 応用編

#### 2-1

`country`に含まれる全ての国と、`city`テーブルに含まれる全ての都市を国と都市を紐づけて取得してください。`city`テーブルに都市が 1 つもない国もありますが、そのような国も出力してください。

:::details **答え**

```sql
SELECT country.Name, city.Name FROM country LEFT JOIN city ON country.Code = city.CountryCode;
```

**出力**

```txt
+----------------------------------------------+-----------------------------------+
| Name                                         | Name                              |
+----------------------------------------------+-----------------------------------+
| Aruba                                        | Oranjestad                        |
| Afghanistan                                  | Kabul                             |
| Afghanistan                                  | Qandahar                          |
| Afghanistan                                  | Herat                             |
| Afghanistan                                  | Mazar-e-Sharif                    |
| Angola                                       | Luanda                            |
| Angola                                       | Huambo                            |
| Angola                                       | Lobito                            |
| Angola                                       | Benguela                          |
| Angola                                       | Namibe                            |
| Anguilla                                     | South Hill                        |
| Anguilla                                     | The Valley                        |
| Albania                                      | Tirana                            |
| Andorra                                      | Andorra la Vella                  |
...省略
...
| Zimbabwe                                     | Gweru                             |
+----------------------------------------------+-----------------------------------+
4086 rows in set (0.00 sec)
```

取得レコード数が正しいか確かめましょう。
:::

#### 2-2

`country`テーブルから、人口の多い順に順位、国名、人口を取得してください。

:::details **答え**

```sql
SELECT RANK() OVER (ORDER BY Population DESC), Name, Population FROM country LIMIT 10;
```

**出力**

```txt
+----------------------------------------+--------------------+------------+
| RANK() OVER (ORDER BY Population DESC) | Name               | Population |
+----------------------------------------+--------------------+------------+
|                                      1 | China              | 1277558000 |
|                                      2 | India              | 1013662000 |
|                                      3 | United States      |  278357000 |
|                                      4 | Indonesia          |  212107000 |
|                                      5 | Brazil             |  170115000 |
|                                      6 | Pakistan           |  156483000 |
|                                      7 | Russian Federation |  146934000 |
|                                      8 | Bangladesh         |  129155000 |
|                                      9 | Japan              |  126714000 |
|                                     10 | Nigeria            |  111506000 |
+----------------------------------------+--------------------+------------+
10 rows in set (0.00 sec)
```

`Rank()`は Window 関数と呼ばれるものの 1 つで、 MySQL では 8.0 から使えるようになりました。
https://dev.mysql.com/doc/refman/8.0/ja/window-functions.html

:::

#### 2-3

`Velbert`という都市がある国の名前、大陸名(`Continent`)、地区名(`Region`)、人口を**1つのクエリで**取得してください。

:::details **答え**

```sql
SELECT Name, Continent, Region, Population FROM country WHERE Code = (SELECT CountryCode FROM city WHERE Name = "Velbert");
```

**出力**

```txt
+---------+-----------+----------------+------------+
| Name    | Continent | Region         | Population |
+---------+-----------+----------------+------------+
| Germany | Europe    | Western Europe |   82164700 |
+---------+-----------+----------------+------------+
1 row in set (0.01 sec)
```

このように 1 つの SQL 文の中にかっこでくくって別の SQL 文を書くような方法をサブクエリと言います。
:::

#### 2-4

話者が多い言語の上位 10 言語の順位(`Rank`)、言語名、話者数合計(`Speakers`)を取得してください。

:::details **答え**

```sql
SELECT 
  RANK() OVER (ORDER BY SUM(countrylanguage.Percentage * country.Population / 100) DESC) AS "Rank", 
  countrylanguage.Language,
  SUM(countrylanguage.Percentage*country.Population/100) AS "Speakers" 
FROM countrylanguage JOIN country ON countrylanguage.CountryCode = country.Code 
GROUP BY countrylanguage.Language 
ORDER BY Speakers DESC LIMIT 10;
```

**出力**

```txt
+------+------------+------------------+
| Rank | Language   | Speakers         |
+------+------------+------------------+
|    1 | Chinese    | 1191843539.22187 |
|    2 | Hindi      |  405633085.47466 |
|    3 | Spanish    |  355029461.90782 |
|    4 | English    |  347077860.65105 |
|    5 | Arabic     |  233839240.44018 |
|    6 | Bengali    |  209304713.12510 |
|    7 | Portuguese |  177595269.43999 |
|    8 | Russian    |  160807559.89702 |
|    9 | Japanese   |  126814106.08493 |
|   10 | Punjabi    |  104025371.70681 |
+------+------------+------------------+
10 rows in set (0.02 sec)
```

割合を掛けているため、人数に小数が出てきます。
:::
