-- solution 1
-- 设置变量，直接select排序后的第N个值
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    set N = N - 1;

  RETURN (
      select
      	ifnull(
        	(select distinct Salary
        	from Employee
        	order by Salary desc limit N, 1), null)
  );
END


-- solution 2
-- select前N个最大的值，再取第N个值，需要判断第N个值是否存在
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
      select if(counts<N, null, mins)
      from  
        (select min(Salary) as mins, count(1) as counts
        from  
            (select distinct Salary from Employee
            order by Salary desc limit N) as a 
        ) as b
  );
END 

