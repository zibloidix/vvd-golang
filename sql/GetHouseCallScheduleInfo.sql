select s.uuid SlotID,
       s.visit_date VisitTime,
       s.duration Duration
  from slots s
  join doctors d 
    on s.doctor_id = d.id 
 where 1=1
   and s.patient_id is null
   and snils = $1 
   and doctor_code = $2
   and s.visit_date > $3
   and s.visit_date < $4