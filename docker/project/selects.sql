-- selects for users
-- сортировка юзеров по имени или телефону
SELECT *
FROM users
ORDER BY name;
SELECT *
FROM users
ORDER BY phone;

-- selects for dialogs, messages, dialog_users
-- подсчет сообщений в диалогах
SELECT
    m.dialog_id,
    count(m.message)
FROM messages AS m
GROUP BY m.dialog_id;
-- вывод списка пользотелей диалога
SELECT
    du.dialog_id,
    du.user_id
FROM dialog_users AS du
GROUP BY du.dialog_id, du.user_id;
-- подсчет сообщений с миенами пользователей
SELECT
    du.dialog_id,
    du.user_id,
    count(m.message),
    (SELECT u.name FROM users AS u WHERE u.id = du.user_id)
FROM dialog_users AS du
JOIN messages AS m ON m.sender_id = du.user_id
GROUP BY du.dialog_id, du.user_id;


-- selects for attachments
-- подсчет типов вложений
SELECT
    at.alias,
    count(a.attachment_type_id)
FROM attachments AS a
JOIN attachment_types AS at ON at.id = a.attachment_type_id
GROUP BY at.alias;
-- вывод по диалогам, типам вложений и подсчет типов
SELECT
    m.dialog_id,
    at.alias,
    count(at.alias)
FROM attachments AS a
JOIN attachment_types AS at ON a.attachment_type_id = at.id
JOIN messages AS m ON m.id = a.message_id
GROUP BY m.dialog_id, at.alias;

-- selects for votes
-- количество опросов в диалоге
SELECT
    m.dialog_id,
    count(v.id)
FROM votes AS v
JOIN messages m ON v.message_id = m.id
GROUP BY m.dialog_id;
-- количество вариантов ответа в опросах
SELECT
    vo.vote_id,
    count(vo.label)
FROM vote_options AS vo
GROUP BY vo.vote_id;
