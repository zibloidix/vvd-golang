select s.rmis_id RmisID,
       s.uuid SlotID,
       s.visit_date VisitTime,
       s.duration Duration
  from slots s
 where uuid = $1