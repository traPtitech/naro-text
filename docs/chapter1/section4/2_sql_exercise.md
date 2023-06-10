# SQL演習問題

基本編・応用編があります。基本編はここまで説明したことだけで解けるはずです。応用編はここまで紹介していないものも使うので、自分で調べて解く必要があります。基本編は解けるようにしましょう。応用編は余裕があったらやってみてください。
SQL を実行するのは Adminer 上でも、`task db`を実行して MySQL にログインしてもどちらでもよいです。自分が好きな方でやってください。

## 基本編

### 1-1

`country`テーブルから、日本(`Name`のカラムが`Japan`)の情報を全て取得してください。

:::details 答え

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

### 1-2

`country`テーブルから、独立年（建国年）(`IndepYear`)が`0`以下の国の情報をすべて取得してください。

:::details 答え

```sql
SELECT * FROM country WHERE IndepYear <= 0;
```

**出力**

```txt
+------+----------+-----------+----------------+-------------+-----------+------------+----------------+------------+------------+--------------+-------------------------+----------------+---------+-------+
| Code | Name     | Continent | Region         | SurfaceArea | IndepYear | Population | LifeExpectancy | GNP        | GNPOld     | LocalName    | GovernmentForm          | HeadOfState    | Capital | Code2 |
+------+----------+-----------+----------------+-------------+-----------+------------+----------------+------------+------------+--------------+-------------------------+----------------+---------+-------+
| CHN  | China    | Asia      | Eastern Asia   |  9572900.00 |     -1523 | 1277558000 |           71.4 |  982268.00 |  917719.00 | Zhongquo     | People'sRepublic        | Jiang Zemin    |    1891 | CN    |
| ETH  | Ethiopia | Africa    | Eastern Africa |  1104300.00 |     -1000 |   62565000 |           45.2 |    6353.00 |    6180.00 | YeItyop´iya  | Republic                | Negasso Gidada |     756 | ET    |
| JPN  | Japan    | Asia      | Eastern Asia   |   377829.00 |      -660 |  126714000 |           80.7 | 3787042.00 | 4192638.00 | Nihon/Nippon | Constitutional Monarchy | Akihito        |    1532 | JP    |
+------+----------+-----------+----------------+-------------+-----------+------------+----------------+------------+------------+--------------+-------------------------+----------------+---------+-------+
3 rows in set (0.59 sec)
```

:::

### 1-3

`country`テーブルから、大陸(`Continent`)がアジア(`Asia`)の国のうち、人口(`Population`)の上位 5 か国の国名(`Name`)と人口を多い順に取得してください。

:::details 答え

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

### 1-4

`country`テーブルから人口が多い順に 11 位から 15 位の国の情報をすべて取得してください。

:::details 答え

```sql
SELECT * FROM country ORDER BY Population DESC LIMIT 5 OFFSET 10;
```

**出力**

```txt
+------+-------------+---------------+-----------------+-------------+-----------+------------+----------------+------------+------------+-------------+----------------------+-------------------------+---------+-------+
| Code | Name        | Continent     | Region          | SurfaceArea | IndepYear | Population | LifeExpectancy | GNP        | GNPOld     | LocalName   | GovernmentForm       | HeadOfState             | Capital | Code2 |
+------+-------------+---------------+-----------------+-------------+-----------+------------+----------------+------------+------------+-------------+----------------------+-------------------------+---------+-------+
| MEX  | Mexico      | North America | Central America |  1958201.00 |      1810 |   98881000 |           71.5 |  414972.00 |  401461.00 | México      | Federal Republic     | Vicente Fox Quesada     |    2515 | MX    |
| DEU  | Germany     | Europe        | Western Europe  |   357022.00 |      1955 |   82164700 |           77.4 | 2133367.00 | 2102826.00 | Deutschland | Federal Republic     | Johannes Rau            |    3068 | DE    |
| VNM  | Vietnam     | Asia          | Southeast Asia  |   331689.00 |      1945 |   79832000 |           69.3 |   21929.00 |   22834.00 | Viêt Nam    | Socialistic Republic | Trân Duc Luong          |    3770 | VN    |
| PHL  | Philippines | Asia          | Southeast Asia  |   300000.00 |      1946 |   75967000 |           67.5 |   65107.00 |   82239.00 | Pilipinas   | Republic             | Gloria Macapagal-Arroyo |     766 | PH    |
| EGY  | Egypt       | Africa        | Northern Africa |  1001449.00 |      1922 |   68470000 |           63.3 |   82710.00 |   75617.00 | Misr        | Republic             | Hosni Mubarak           |     608 | EG    |
+------+-------------+---------------+-----------------+-------------+-----------+------------+----------------+------------+------------+-------------+----------------------+-------------------------+---------+-------+
5 rows in set (0.00 sec)
```

:::

### 1-5

`countrylanguage`テーブルから、言語(`Language`)に日本語(`Japanese`)が含まれる国を数えて、数を`Japanese`という項目名で取得してください。

:::details 答え

```sql
SELECT COUNT(*) AS Japanese FROM countrylanguage WHERE Language = "Japanese";
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

### 1-6

`country`テーブルから、大陸(`Continent`)ごとの国の数を数えて大陸ごとに取得してください。

:::details 答え

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

### 1-7

`city`テーブルと`country`テーブルを使って、国名と首都の一覧を、国名を`Country`、首都を`Capital`として表示してください。

ヒント：`country`テーブルの`Capital`カラムは、その国の首都の`city`テーブルでの`ID`を指しています。

:::details 答え

```sql
SELECT country.Name AS Country, city.Name AS Capital FROM country JOIN city ON country.Capital = city.ID;
```

**出力**

```txt
+---------------------------------------+------------------------------------+
| Country                               | Capital                            |
+---------------------------------------+------------------------------------+
| Aruba                                 | Oranjestad                         |
| Afghanistan                           | Kabul                              |
| Angola                                | Luanda                             |
| Anguilla                              | The Valley                         |
| Albania                               | Tirana                             |
| Andorra                               | Andorra la Vella                   |
| Netherlands Antilles                  | Willemstad                         |
| United Arab Emirates                  | Abu Dhabi                          |
| Argentina                             | Buenos Aires                       |
| Armenia                               | Yerevan                            |
| American Samoa                        | Fagatogo                           |
| Antigua and Barbuda                   | Saint John´s                       |
| Australia                             | Canberra                           |
| Austria                               | Wien                               |
| Azerbaijan                            | Baku                               |
| Burundi                               | Bujumbura                          |
省略...
```

:::

### 1-8

`country`テーブルから、大陸(`Continent`)が北アメリカ(`North America`)または南アメリカ(`South America`)の国の情報をすべて取得してください。この問題は答えが 2 通りあります。余裕のある人は 2 つ考えてみましょう。

:::details 答え

```sql
SELECT * FROM country WHERE Continent = "North America" OR Continent = "South America";
--または
SELECT * FROM country WHERE Continent IN ("North America", "South America");
```

**出力**

```txt
+------+----------------------------------+---------------+-----------------+-------------+-----------+------------+----------------+------------+------------+-------------------------------------+----------------------------------------------+--------------------------------------+---------+-------+
| Code | Name                             | Continent     | Region          | SurfaceArea | IndepYear | Population | LifeExpectancy | GNP        | GNPOld     | LocalName                           | GovernmentForm                               | HeadOfState                          | Capital | Code2 |
+------+----------------------------------+---------------+-----------------+-------------+-----------+------------+----------------+------------+------------+-------------------------------------+----------------------------------------------+--------------------------------------+---------+-------+
| ABW  | Aruba                            | North America | Caribbean       |      193.00 |      NULL |     103000 |           78.4 |     828.00 |     793.00 | Aruba                               | Nonmetropolitan Territory of The Netherlands | Beatrix                              |     129 | AW    |
| AIA  | Anguilla                         | North America | Caribbean       |       96.00 |      NULL |       8000 |           76.1 |      63.20 |       NULL | Anguilla                            | Dependent Territory of the UK                | Elisabeth II                         |      62 | AI    |
| ANT  | Netherlands Antilles             | North America | Caribbean       |      800.00 |      NULL |     217000 |           74.7 |    1941.00 |       NULL | Nederlandse Antillen                | Nonmetropolitan Territory of The Netherlands | Beatrix                              |      33 | AN    |
| ARG  | Argentina                        | South America | South America   |  2780400.00 |      1816 |   37032000 |           75.1 |  340238.00 |  323310.00 | Argentina                           | Federal Republic                             | Fernando de la Rúa                   |      69 | AR    |
| ATG  | Antigua and Barbuda              | North America | Caribbean       |      442.00 |      1981 |      68000 |           70.5 |     612.00 |     584.00 | Antigua and Barbuda                 | Constitutional Monarchy                      | Elisabeth II                         |      63 | AG    |
| BHS  | Bahamas                          | North America | Caribbean       |    13878.00 |      1973 |     307000 |           71.1 |    3527.00 |    3347.00 | The Bahamas                         | Constitutional Monarchy                      | Elisabeth II                         |     148 | BS    |                      
...省略
```

:::

### 1-9

`country`テーブルに以下の情報を持った国を挿入してください。

- 国名 `traP`
- 国コード `TRP`
- 大陸 `Asia`
- 地域 `Eastern Asia`
- 現地での呼び名(`LocalName`) `traP`
- 政治体制(`GovernmentForm`) 立憲君主制(`Constitutional Monarchy`)
- 人口 `500`人
- 2 文字国コード(`Code2`) `TP`

:::details 答え

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

### 1-10

1-9 で作った国「`traP`」の独立年(`IndepYear`)を 2015 年にしましょう。

:::details 答え

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

### 1-11

`country`に追加した国「`traP`」を削除しましょう。

:::details 答え

```sql
DELETE FROM country WHERE Code = "TRP";
```

**確認**

```sql
SELECT * FROM country WHERE Code = "TRP";
```

**出力**

```txt
Empty set (0.00 sec)
```

:::

基本問題が解けたら、一番難しかった問題を講習会チャンネルに投稿しましょう。

## 応用編

### 2-1

`country`に含まれる全ての国と、`city`テーブルに含まれる全ての都市を国と都市を紐づけて取得してください。 **`city`テーブルに都市が 1 つもない国もありますが、そのような国も出力してください。** 全部で 4086 行になるはずです。

:::details 答え

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

### 2-2

`country`テーブルから、人口の多い順に順位(`Rank`)、国名、人口を取得してください。

:::details 答え

```sql
SELECT RANK() OVER (ORDER BY Population DESC) AS Rank, Name, Population FROM country LIMIT 10;
```

**出力**

```txt
+------+--------------------+------------+
| Rank | Name               | Population |
+------+--------------------+------------+
|    1 | China              | 1277558000 |
|    2 | India              | 1013662000 |
|    3 | United States      |  278357000 |
|    4 | Indonesia          |  212107000 |
|    5 | Brazil             |  170115000 |
|    6 | Pakistan           |  156483000 |
|    7 | Russian Federation |  146934000 |
|    8 | Bangladesh         |  129155000 |
|    9 | Japan              |  126714000 |
|   10 | Nigeria            |  111506000 |
+------+--------------------+------------+
10 rows in set (0.00 sec)
```

`Rank()`は Window 関数と呼ばれるものの 1 つで、 MySQL では 8.0 から使えるようになりました。

公式ドキュメント [https://dev.mysql.com/doc/refman/8.0/ja/window-functions.html](https://dev.mysql.com/doc/refman/8.0/ja/window-functions.html)

:::

### 2-3

`Velbert`という都市がある国の名前、大陸名(`Continent`)、地区名(`Region`)、人口を**1つのクエリで**取得してください。

:::details 答え

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

### 2-4

話者が多い言語の上位 10 言語の順位(`Rank`)、言語名、話者数合計(`Speakers`)を取得してください。

:::details 答え

```sql
SELECT 
  RANK() OVER (ORDER BY SUM(countrylanguage.Percentage * country.Population / 100) DESC) AS Rank, 
  countrylanguage.Language,
  SUM(countrylanguage.Percentage*country.Population/100) AS Speakers 
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

応用問題が解けたら、一番難しかった問題を講習会チャンネルに投稿しましょう。
