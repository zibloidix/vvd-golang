 
select rmis_id RmisID
  from slots s
  where 1=1
    and s.rmis_id = $1
    and s.patient_id is not null