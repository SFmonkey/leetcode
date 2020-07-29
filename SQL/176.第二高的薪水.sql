# Solution 1
select max(Salary) as "SecondHighestSalary"
from Employee
where
	Salary != ( 
		select max(Salary) from Employee
	)
;


# Solution 2
select 
ifnull(
	(select Salary 
	 from Employee 
	 order by Salary limit 1, 1),
	null 
)
;
