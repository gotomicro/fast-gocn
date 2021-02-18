create table topic_cate
(
    id         int auto_increment
        primary key,
    title      varchar(255) not null,
    node_id    int          not null,
    created_at datetime     null,
    updated_at datetime     null,
    deleted_at datetime     null,
    cnt_topic  int          not null
);

create index idx_topic_cate_node_id
    on topic_cate (node_id);

INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (1, '每日新闻', 18, '2021-01-05 13:52:30', '2021-02-02 10:58:01', null, 1330);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (2, '分享', 52, '2021-01-05 13:52:35', '2021-02-01 15:59:10', null, 97);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (3, 'TiDB', 46, '2021-01-05 13:52:50', '2021-02-02 11:30:50', null, 139);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (4, 'Docker', 12, '2021-01-05 13:52:55', '2021-01-20 15:02:43', null, 136);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (5, 'Java', 20, '2021-01-05 13:53:00', '2021-01-19 11:11:22', null, 38);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (6, '闲聊', 53, '2021-01-05 13:53:06', '2021-01-25 18:38:15', null, 21);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (7, '原创分享', 3, '2021-01-05 13:53:16', '2021-02-01 15:26:47', null, 317);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (8, '招聘', 25, '2021-01-05 13:53:21', '2021-01-28 15:34:08', null, 446);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (9, '开源推荐', 2, '2021-01-05 13:53:26', '2021-01-24 16:37:58', null, 104);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (10, '资料分享', 9, '2021-01-05 13:53:31', '2021-01-12 01:57:17', null, 46);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (11, 'GoFrame', 14, '2021-01-05 13:54:27', '2021-01-12 03:15:49', null, 14);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (12, 'Go问答', 1, '2021-01-05 13:54:32', '2021-01-29 15:40:09', null, 2963);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (13, '微服务', 50, '2021-01-05 13:54:37', '2021-01-25 19:10:35', null, 53);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (14, 'Kubernetes', 11, '2021-01-05 13:55:27', '2021-01-29 17:49:21', null, 93);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (15, 'Nightingale', 58, '2021-01-05 13:56:59', '2021-01-05 14:49:24', null, 3);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (16, 'Meetup', 28, '2021-01-05 13:58:05', '2021-01-12 04:46:10', null, 24);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (17, '开发工具', 4, '2021-01-05 13:58:55', '2021-01-23 17:08:28', null, 20);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (18, '译文', 39, '2021-01-05 13:59:46', '2021-01-12 01:56:57', null, 26);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (19, '线上活动', 40, '2021-01-05 13:59:51', '2021-01-12 06:25:39', null, 18);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (20, 'Gin', 13, '2021-01-05 14:07:23', '2021-01-12 05:52:17', null, 15);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (21, 'GopherChina', 27, '2021-01-05 14:08:39', '2021-01-12 06:35:11', null, 21);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (22, '区块链', 10, '2021-01-05 14:09:05', '2021-01-11 23:41:12', null, 2);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (23, '云原生', 55, '2021-01-05 14:23:17', '2021-01-05 15:30:45', null, 32);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (24, 'Beego', 15, '2021-01-05 14:24:58', '2021-01-12 06:34:21', null, 232);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (25, '代码分享', 8, '2021-01-05 14:25:59', '2021-01-12 01:57:58', null, 22);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (26, 'GORM', 44, '2021-01-05 14:29:02', '2021-01-12 06:20:14', null, 18);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (27, '求职', 16, '2021-01-05 14:29:48', '2021-01-05 14:40:06', null, 4);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (28, 'dubbogo', 57, '2021-01-05 14:33:26', '2021-01-18 18:11:00', null, 9);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (29, 'GoChassis', 59, '2021-01-05 14:39:26', '2021-01-05 14:41:37', null, 3);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (30, 'TarsGo', 51, '2021-01-05 14:48:18', '2021-01-05 14:52:32', null, 3);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (31, 'Python', 19, '2021-01-05 15:00:54', '2021-01-05 15:29:34', null, 9);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (32, 'Go夜读', 48, '2021-01-05 15:09:52', '2021-01-12 04:08:10', null, 90);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (33, 'Mr.Go会客厅', 56, '2021-01-05 15:12:34', '2021-01-05 15:20:11', null, 2);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (34, 'Node.js', 22, '2021-01-05 15:22:43', '2021-01-05 15:22:43', null, 1);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (35, '反馈', 24, '2021-01-05 15:25:51', '2021-01-11 23:43:39', null, 7);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (36, 'Coding', 6, '2021-01-05 15:34:39', '2021-01-05 15:34:39', null, 1);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (37, 'Prometheus', 32, '2021-01-05 15:34:44', '2021-01-12 01:08:56', null, 4);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (38, 'gRPC', 34, '2021-01-05 15:34:59', '2021-01-05 15:34:59', null, 1);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (39, '书籍', 54, '2021-01-05 15:35:25', '2021-01-05 15:35:25', null, 1);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (40, 'Gitea', 45, '2021-01-05 15:42:37', '2021-01-12 05:46:07', null, 15);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (41, 'NoPoint', 26, '2021-01-05 15:55:59', '2021-01-05 15:55:59', null, 1);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (42, 'LiteIDE', 47, '2021-01-05 15:56:09', '2021-01-12 06:09:46', null, 15);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (43, '算法', 5, '2021-01-05 15:59:53', '2021-01-11 23:37:43', null, 3);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (44, '公告', 23, '2021-01-11 23:40:11', '2021-01-11 23:55:54', null, 2);
INSERT INTO fastgocn.topic_cate (id, title, node_id, created_at, updated_at, deleted_at, cnt_topic) VALUES (45, 'Hackathon', 29, '2021-01-11 23:47:52', '2021-01-12 04:51:59', null, 4);