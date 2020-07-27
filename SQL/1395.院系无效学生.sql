# Solution 1
select s.id, s.name
from Students s left join Departments d on s.department_id = d.id
where 
    d.id is null;


# Solution 2
select s.id, s.name
from Students s
where 
	s.department_id not in (
		select id from Departments 
	)
