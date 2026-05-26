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
    `id`                BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `node_url`          VARCHAR(255) NOT NULL COMMENT '合约发布网站地址，如127.0.0.1:8454',
    `chain_id`          varchar(20)           DEFAULT '56' COMMENT '链 ID: 56=BSC 主网 97=BSC 测试网',
    `contract_address`  VARCHAR(255) NOT NULL COMMENT '合约地址',
    `contract_abi`      json         NOT NULL COMMENT '合约ABI',
    `contract_bin`      json         NOT NULL COMMENT '合约BIN',
    `publisher_address` VARCHAR(255) NOT NULL COMMENT '合约发布者地址',
    `created_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='合约表';

CREATE TABLE `poolbases`
(
    `id`                       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
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
    `state`                    varchar(100) DEFAULT NULL COMMENT '池子状态: 0未开启 1进行中 2已结算 3清算中 4未开启等',
    `lend_debt_token`          varchar(100) DEFAULT NULL COMMENT '出借侧池子代币/合约地址(出借凭证)',
    `borrow_debt_token`          varchar(100) DEFAULT NULL COMMENT '借入(抵押)侧池子代币/合约地址(质押资产)',
    `auto_liquidate_threshold` varchar(100) DEFAULT NULL COMMENT '自动清算阈值(精度值)',
    `borrow_token_info`        json         DEFAULT NULL COMMENT '借入(抵押)代币信息: tokenName, tokenLogo, tokenPrice, borrowFee 等',
    `lend_token_info`          json         DEFAULT NULL COMMENT '出借代币信息: tokenName, tokenLogo, tokenPrice, lendFee 等',
    `chain_id`                 varchar(20)  DEFAULT '56' COMMENT '链 ID: 56=BSC 主网 97=BSC 测试网',
    `lend_token_symbol`        varchar(100) DEFAULT NULL COMMENT '出借代币符号，如 BUSD',
    `borrow_token_symbol`      varchar(100) DEFAULT NULL COMMENT '借入(抵押)代币符号，如 BTC',
    `created_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='借贷池主表(按链+pool_id)';

CREATE TABLE `pooldata`
(
    `id`                       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `chain_id`                 varchar(20)  DEFAULT '56' COMMENT '链 ID，与 poolbases 一致',
    `pool_id`                  varchar(50)  DEFAULT NULL COMMENT '池子 ID，与 poolbases.pool_id 对应',
    `settle_amount_lend`       varchar(100) DEFAULT NULL COMMENT '结算时借出侧(贷方)金额',
    `settle_amount_borrow`     varchar(100) DEFAULT NULL COMMENT '结算时借入侧(抵押)金额',
    `finish_amount_lend`       varchar(100) DEFAULT NULL COMMENT '已完成/归还的借出侧金额',
    `finish_amount_borrow`     varchar(100) DEFAULT NULL COMMENT '已完成/归还的借入侧金额',
    `liquidation_amoun_lend`   varchar(100) DEFAULT NULL COMMENT '清算产生的借出侧金额',
    `liquidation_amoun_borrow` varchar(100) DEFAULT NULL COMMENT '清算产生的借入侧金额',
    `created_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='池子结算与清算数据表';


CREATE TABLE `token_info`
(
    `id`           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `symbol`       varchar(100) DEFAULT NULL COMMENT '代币符号，如 BUSD/BTC',
    `logo`         varchar(150) DEFAULT NULL COMMENT '代币 logo URL',
    `price`        varchar(50)  DEFAULT NULL COMMENT '价格(精度值，用于估值与清算)',
    `token`        varchar(100) DEFAULT NULL COMMENT '代币合约地址',
    `chain_id`     varchar(20)  DEFAULT '56' COMMENT '链 ID: 56=BSC 97=测试网',
    `contract_abi` json NOT NULL COMMENT '代币ABI',
    `contract_bin` json NOT NULL COMMENT '代币BIN',
    `decimals`     int(11) NOT NULL COMMENT '代币精度(小数位数)',
    `created_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`        DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='代币信息表(按链)';


