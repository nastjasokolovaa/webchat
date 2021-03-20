-- представление активности пользователей
CREATE VIEW active_users (id, name, count_messages)
    AS SELECT
           u.id,
           u.name,
           count(m.message)
       FROM messages AS m
       JOIN users AS u ON u.id = m.sender_id
       GROUP BY u.id, u.name
       ORDER BY count(m.message) DESC;

-- представление по выводу и подсчету популярных типов вложения
CREATE VIEW popular_attachment (dialog_id, attachment_type, count_types)
AS SELECT
       m.dialog_id,
       at.alias,
       count(at.alias)
   FROM attachments AS a
   JOIN attachment_types AS at ON a.attachment_type_id = at.id
   JOIN messages AS m ON m.id = a.message_id
   GROUP BY m.dialog_id, at.alias;
