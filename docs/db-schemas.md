## Users

- id
- username
- password: hash
- name
- role: student | teacher | admin

## Classes

- id
- name
- teachers: [userId]
- students: [userId]
- subjects: [constants]
- startDate
- endDate

## Subjects

- id
- name
