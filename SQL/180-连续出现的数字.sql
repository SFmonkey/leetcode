-- 借助窗口函数lead()实现列平移
select distinct Num as "ConsecutiveNums"
from 
    (
        select Num, lead(Num, 1) over() as Num1, lead(Num, 2) over() as Num2
        from Logs
    ) as tmp
where tmp.Num = tmp.Num1 and tmp.Num1 = tmp.Num2
