-- ユーザー情報
CREATE TABLE `users` (
  `user_id` INT PRIMARY KEY,
  `name` VARCHAR(50) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 認証情報
CREATE TABLE `user_auths` (
  `id` INT PRIMARY KEY AUTO_INCREMENT,
  `user_id` INT NOT NULL,
  `identity_type` VARCHAR(50) NOT NULL,
  `identifier` VARCHAR(100) NOT NULL,
  `credential` VARCHAR(100) NOT NULL,
  UNIQUE KEY `user_id_identity_type` (`user_id`, `identity_type`)
)
;

-- jwt tokenのブラックリスト
CREATE TABLE `diasabled_tokens` (
  `user_id` INT PRIMARY KEY AUTO_INCREMENT,
  `token` VARCHAR(100) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 初期ユーザー
INSERT INTO `users` (`user_id`, `name`) VALUES (1, 'admin');

-- 初期ユーザーの認証情報
INSERT INTO `user_auths`
  (`user_id`, `identity_type`, `identifier`, `credential`)
VALUES
  (1, 'email', 'mail@example.com', 'password');
