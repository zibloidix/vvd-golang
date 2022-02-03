
select 1001 PatientID,
       h.id			MOID,
       h.oid		MOOID,
       h.name		MOName,
       h.address	MOAddress,
       h.phone		MOPhone,
       d.snils || '.' || d.doctor_code		ResourceID,
       d.name  || ' (' || s.name  || ')' 	ResourceName
  from hospitals h
  join doctors d
    on d.hospital_id = h.id
  join specs s
    on d.spec_id = s.id
  join districts ds
    on ds.hospital_id = h.id
 where 1=1
   and ds.city = $1
   and ds.street like '%' || $2 || '%'
   and ds.house = $3