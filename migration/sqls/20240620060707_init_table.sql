-- +goose Up
-- +goose StatementBegin
-- ユーザーの作成
CREATE USER 'application' @'%' IDENTIFIED BY 'app_password';
CREATE USER 'ro_user' @'%' IDENTIFIED BY 'ro_password';
CREATE USER 'rw_user' @'%' IDENTIFIED BY 'rw_password';
-- 権限設定
GRANT SELECT,
    INSERT,
    UPDATE,
    DELETE ON app.* TO 'cd_user' @'%';
GRANT SELECT,
    INSERT,
    UPDATE,
    DELETE ON app.* TO 'application' @'%';
GRANT SELECT,
    INSERT,
    UPDATE,
    DELETE ON app.* TO 'rw_user' @'%';
GRANT SELECT ON app.* TO 'ro_user' @'%';
-- 権限の適用
FLUSH PRIVILEGES;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP USER 'ro_user' @'%';
DROP USER 'rw_user' @'%';
DROP USER 'application' @'%';
-- 権限の適用
FLUSH PRIVILEGES;
-- +goose StatementEnd