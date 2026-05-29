CREATE TABLE `user`
(
    `id`       bigint NOT NULL AUTO_INCREMENT,
    `login`    varchar(255) DEFAULT NULL COMMENT '用户账号',
    `nike`     varchar(255) DEFAULT NULL COMMENT '用户昵称',
    `password` varchar(255) DEFAULT NULL COMMENT '加密后的密码',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE `contract`
(
    `id`                BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `node_url`          VARCHAR(255) NOT NULL COMMENT '合约发布网站地址，如127.0.0.1:8454',
    `chain_id`          varchar(20)  NOT NULL COMMENT '链 ID: 56=BSC 主网 97=BSC 测试网',
    `contract_address`  VARCHAR(255) NOT NULL COMMENT '合约地址',
    `contract_name`    varchar(100) COMMENT '合约名称',
    `tx_hash`  varchar(100) COMMENT 'txHash',
    `publisher_address` VARCHAR(255) NOT NULL COMMENT '合约发布者地址',
    `is_token`          tinyint(1) DEFAULT 0 COMMENT '是否为代币合约',
    `token_symbol`      varchar(50) DEFAULT NULL COMMENT '代币符号',
    `token_decimals`    int(11) DEFAULT 0 COMMENT '代币精度(小数点位数)',
    `created_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合约表';

CREATE TABLE `poolbases`
(
    `id`                       BIGINT  NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `contract_id`              BIGINT  DEFAULT 0 COMMENT '关联的合约 ID，对应 contract.id',
    `pool_id`                  int(11) DEFAULT NULL COMMENT '业务池子 ID，与链上 pool 对应',
    `settle_time`              varchar(100) DEFAULT NULL COMMENT '结算开始时间戳',
    `end_time`                 varchar(100) DEFAULT NULL COMMENT '池子结束时间戳',
    `interest_rate`            varchar(100) DEFAULT NULL COMMENT '利率(精度值，如 10000000 表示 1%)',
    `max_supply`               varchar(100) DEFAULT NULL COMMENT '池子最大可借/可存额度(最小单位)',
    `lend_supply`              varchar(100) DEFAULT NULL COMMENT '当前已出借总量',
    `borrow_supply`            varchar(100) DEFAULT NULL COMMENT '当前已借入(抵押)总量',
    `mortgage_rate`            varchar(100) DEFAULT NULL COMMENT '抵押率(精度值，如 10000000=100%)',
    `lend_token`               varchar(100) DEFAULT NULL COMMENT '出借资产合约地址(用户借出的币)',
    `borrow_token`             varchar(100) DEFAULT NULL COMMENT '借入(抵押)资产合约地址(用户抵押的币)',
    `state`                    varchar(100) DEFAULT NULL COMMENT '池子状态: MATCH=撮合中 EXECUTION=执行期 FINISH=结束 LIQUIDATION=清算 UNDONE=未成立',
    `lend_debt_token`          varchar(100) DEFAULT NULL COMMENT '出借侧池子代币/合约地址(出借凭证)',
    `borrow_debt_token`          varchar(100) DEFAULT NULL COMMENT '借入(抵押)侧池子代币/合约地址(质押资产)',
    `auto_liquidate_threshold` varchar(100) DEFAULT NULL COMMENT '自动清算阈值(精度值)',
    `chain_id`                 varchar(20)  DEFAULT '56' COMMENT '链 ID: 56=BSC 主网 97=BSC 测试网',
    `settle_amount_lend`       varchar(100) DEFAULT NULL COMMENT '结算时借出侧(贷方)金额',
    `settle_amount_borrow`     varchar(100) DEFAULT NULL COMMENT '结算时借入侧(抵押)金额',
    `finish_amount_lend`       varchar(100) DEFAULT NULL COMMENT '已完成/归还的借出侧金额',
    `finish_amount_borrow`     varchar(100) DEFAULT NULL COMMENT '已完成/归还的借入侧金额',
    `liquidation_amoun_lend`   varchar(100) DEFAULT NULL COMMENT '清算产生的借出侧金额',
    `liquidation_amoun_borrow` varchar(100) DEFAULT NULL COMMENT '清算产生的借入侧金额',
    `created_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='借贷池主表';
