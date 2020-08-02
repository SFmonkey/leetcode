-- 从表2中找到有多少大于表1中数值的记录，就是排名
-- count(distinct)记录排名，不用group by的话只返回最后一条结果
select s1.Score as "Score",
count(distinct(s2.Score)) as "Rank"
from Scores as s1, Scores as s2
where s1.score <= s2.score
group by s1.id
order by s1.score desc;
