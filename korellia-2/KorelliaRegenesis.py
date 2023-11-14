import copy
import json
import sys


# fix_amounts takes a coins array divides the tkyve entry by 1000 and removes all other entries.
# This function keeps the reference of the original array.
def fix_cosmos_amounts(amounts):
    tkyve_amount = {"denom": "tkyve", "amount": "0"}
    for i in range(len(amounts)):
        if amounts[0]["denom"] == "tkyve":
            tkyve_amount["amount"] = str(int(float(amounts[0]["amount"]) / 1000))
        del amounts[0]
    amounts.append(tkyve_amount)


def fix_plain_attribute(entry, attribute_name):
    entry[attribute_name] = str(int(float(entry[attribute_name]) / 1000))


# Returns a default genesis file with correct params
def initialiseDefaultGenesis():
    genesis = dict()
    genesis["genesis_time"] = "2023-11-10T17:30:00.000000Z"
    genesis["chain_id"] = "korellia-2"
    genesis["initial_height"] = "8589000"
    genesis["consensus_params"] = {
        "block": {
            "max_bytes": "22020096",
            "max_gas": "10000000000"  # 10,000,000,000
        },
        "evidence": {
            "max_age_num_blocks": "100000",
            "max_age_duration": "172800000000000",
            "max_bytes": "1048576"
        },
        "validator": {
            "pub_key_types": [
                "ed25519"
            ]
        },
        "version": {
            "app": "0"
        }
    }
    genesis["app_hash"] = ""
    genesis["app_state"] = dict()

    # 06-solomachine
    genesis["app_state"]["06-solomachine"] = None

    # 07-tendermint
    genesis["app_state"]["07-tendermint"] = None

    # auth -> will be injected from Korellia
    genesis["app_state"]["auth"] = dict()

    # authz
    genesis["app_state"]["authz"] = {"authorization": []}

    # bank -> will be injected from Korellia
    genesis["app_state"]["bank"] = dict()

    # bundles -> will be injected from Korellia
    genesis["app_state"]["bundles"] = dict()

    # capability
    genesis["app_state"]["capability"] = {"index": "1", "owners": []}

    # consensus
    genesis["app_state"]["consensus"] = None

    # crisis
    genesis["app_state"]["crisis"] = {
        "constant_fee": {
            "denom": "tkyve",
            "amount": "1000000000"  # 1,000 $KYVE
        }
    }

    # delegation -> will be injected from Korellia
    genesis["app_state"]["delegation"] = dict()

    # distribution -> will be injected from Korellia
    genesis["app_state"]["distribution"] = dict()

    # evidence
    genesis["app_state"]["evidence"] = {"evidence": []}

    # feegrant
    genesis["app_state"]["feegrant"] = {"allowances": []}

    # feeibc
    genesis["app_state"]["feeibc"] = {
        "identified_fees": [],
        "fee_enabled_channels": [],
        "registered_payees": [],
        "registered_counterparty_payees": [],
        "forward_relayers": []
    }

    # funders -< will be injected during gentx handling
    genesis["app_state"]["funders"] = {
        "params": {
            "min_funding_amount": "1000000000",
            "min_funding_amount_per_bundle": "100000",
            "min_funding_multiple": "20"
        },
        "funder_list": [],
        "funding_list": [],
        "funding_state_list": []
    }

    # genutil -> will be injected during gentx handling
    genesis["app_state"]["genutil"] = {
      "gen_txs": []
    }

    # global
    genesis["app_state"]["global"] = {
        "params": {
            "min_gas_price": "0.020000000000000000",
            "burn_ratio": "0.000000000000000000",
            "gas_adjustments": [
                {"amount": "50000000", "type": "/cosmos.staking.v1beta1.MsgCreateValidator"},
                {"amount": "50000000", "type": "/kyve.stakers.v1beta1.MsgCreateStaker"},
                {"amount": "50000000", "type": "/kyve.funders.v1beta1.MsgCreateFunder"}
            ],
            "gas_refunds": [],
            "min_initial_deposit_ratio": "0.000000000000000000"  # deprecated
        }
    }

    # gov -> will be injected from Korellia
    genesis["app_state"]["gov"] = dict()

    # group
    genesis["app_state"]["group"] = {
        "group_seq": "0",
        "groups": [],
        "group_members": [],
        "group_policy_seq": "0",
        "group_policies": [],
        "proposal_seq": "0",
        "proposals": [],
        "votes": []
    }

    genesis["app_state"]["ibc"] = {
        "client_genesis": {
            "clients": [],
            "clients_consensus": [],
            "clients_metadata": [],
            "params": {
                "allowed_clients": [
                    "06-solomachine",
                    "07-tendermint",
                    "09-localhost"
                ]
            },
            "create_localhost": False,
            "next_client_sequence": "0"
        },
        "connection_genesis": {
            "connections": [],
            "client_connection_paths": [],
            "next_connection_sequence": "0",
            "params": {
                "max_expected_time_per_block": "30000000000"
            }
        },
        "channel_genesis": {
            "channels": [],
            "acknowledgements": [],
            "commitments": [],
            "receipts": [],
            "send_sequences": [],
            "recv_sequences": [],
            "ack_sequences": [],
            "next_channel_sequence": "0"
        }
    }

    genesis["app_state"]["interchainaccounts"] = {
        "controller_genesis_state": {
            "active_channels": [],
            "interchain_accounts": [],
            "params": {
                "controller_enabled": True
            },
            "ports": []
        },
        "host_genesis_state": {
            "active_channels": [],
            "interchain_accounts": [],
            "params": {
                "allow_messages": [
                    "*"
                ],
                "host_enabled": True
            },
            "port": "icahost"
        }
    }

    genesis["app_state"]["mint"] = {
        "minter": {
            "annual_provisions": "0.000000000000000000",
            "inflation": "0.050000000000000000"
        },
        "params": {
            "blocks_per_year": "6311520",
            "goal_bonded": "0.670000000000000000",
            "inflation_max": "0.050000000000000000",
            "inflation_min": "0.050000000000000000",
            "inflation_rate_change": "0.130000000000000000",
            "mint_denom": "tkyve"
        }
    }

    genesis["app_state"]["packetfowardmiddleware"] = {
        "in_flight_packets": {},
        "params": {
            "fee_percentage": "0.000000000000000000"
        }
    }

    genesis["app_state"]["params"] = None

    # pool -> will be injected from Korellia
    genesis["app_state"]["pool"] = dict()

    genesis["app_state"]["query"] = None

    genesis["app_state"]["slashing"] = {
        "params": {
            "signed_blocks_window": "100",
            "min_signed_per_window": "0.500000000000000000",
            "downtime_jail_duration": "600s",
            "slash_fraction_double_sign": "0.050000000000000000",
            "slash_fraction_downtime": "0.010000000000000000"
        },
        "signing_infos": [],
        "missed_blocks": []
    }

    # stakers -> will be injected from Korellia
    genesis["app_state"]["stakers"] = dict()

    # staking -> will be injected from Korellia
    genesis["app_state"]["staking"] = dict()

    genesis["app_state"]["team"] = {
      "authority": {
        "total_rewards": "0",
        "rewards_claimed": "0"
      },
      "account_list": [],
      "account_count": "0"
    }

    genesis["app_state"]["transfer"] = {
        "port_id": "transfer",
        "denom_traces": [],
        "params": {
            "send_enabled": True,
            "receive_enabled": True
        },
        "total_escrowed": []
    }

    genesis["app_state"]["upgrade"] = {}
    genesis["app_state"]["vesting"] = {}

    return genesis


def migrateKorelliaToNewGenesis(new, old):
    # auth -> copy all existing accounts
    new["app_state"]["auth"] = copy.deepcopy(old["app_state"]["auth"])
    for entry in new["app_state"]["auth"]["accounts"]:
        # There are 3 types of accounts: Base,Module,Vesting
        if entry["@type"] == "/cosmos.auth.v1beta1.BaseAccount":
            entry["account_number"] = "0"
            entry["pub_key"] = None
            entry["sequence"] = "0"
        elif entry["@type"] == "/cosmos.auth.v1beta1.ModuleAccount":
            entry["base_account"]["account_number"] = "0"
            entry["base_account"]["pub_key"] = None
            entry["base_account"]["sequence"] = "0"
        else:
            # Convert Vesting accounts to normal base accounts
            clone_entry = copy.deepcopy(entry)
            entry.clear()
            entry["@type"] = "/cosmos.auth.v1beta1.BaseAccount"
            entry["account_number"] = "0"
            entry["address"] = clone_entry["base_vesting_account"]["base_account"]["address"]
            entry["pub_key"] = None
            entry["sequence"] = "0"

    # bank -> divide every balance by 1000, to have 6 decimals instead of 9
    total_supply = 0
    converted_bank_balances = copy.deepcopy(old["app_state"]["bank"]["balances"])
    for entry in converted_bank_balances:
        fix_cosmos_amounts(entry["coins"])
        total_supply += int(entry["coins"][0]["amount"])

    new["app_state"]["bank"] = {
      "params": {
        "send_enabled": [],
        "default_send_enabled": True
      },
      "balances": converted_bank_balances,
      "supply": [{'amount': str(total_supply), 'denom': 'tkyve'}],
      "denom_metadata": [],
      "send_enabled": []
    }

    # bundles -> keep finalized bundles, reset current progress
    converted_finalized_bundles = copy.deepcopy(old["app_state"]["bundles"]["finalized_bundle_list"])
    for entry in converted_finalized_bundles:
        entry["stake_security"] = {
            "valid_vote_power": None,
            "total_vote_power": None
        }

    new["app_state"]["bundles"] = {
      "params": {
        "upload_timeout": "600",
        "storage_cost": "0.103996000000000000",
        "network_fee": "0.010000000000000000",
        "max_points": "360"
      },
      "bundle_proposal_list": [],
      "finalized_bundle_list": converted_finalized_bundles,
      "round_robin_progress_list": [],
      "bundle_version_map": {
        "versions": [{
            "height": "8589000",
            "version": "2"
        }]
      }
    }

    # delegation -> reduce from 9 to 6 decimals
    converted_delegator_list = copy.deepcopy(old["app_state"]["delegation"]["delegator_list"])
    for entry in converted_delegator_list:
        fix_plain_attribute(entry, "initial_amount")

    converted_delegation_data_list = copy.deepcopy(old["app_state"]["delegation"]["delegation_data_list"])
    for entry in converted_delegation_data_list:
        fix_plain_attribute(entry, "current_rewards")
        fix_plain_attribute(entry, "total_delegation")

    undelegation_queue_entry_list = copy.deepcopy(old["app_state"]["delegation"]["undelegation_queue_entry_list"])
    for entry in undelegation_queue_entry_list:
        fix_plain_attribute(entry, "amount")

    new["app_state"]["delegation"] = {
      "params": {
        "unbonding_delegation_time": "3600",
        "redelegation_cooldown": "3600",
        "redelegation_max_amount": "5",
        "vote_slash": "0.020000000000000000",
        "upload_slash": "0.050000000000000000",
        "timeout_slash": "0.010000000000000000"
      },
      "delegator_list": converted_delegator_list,
      "delegation_entry_list": copy.deepcopy(old["app_state"]["delegation"]["delegation_entry_list"]),
      "delegation_data_list": converted_delegation_data_list,
      "delegation_slash_list": copy.deepcopy(old["app_state"]["delegation"]["delegation_slash_list"]),
      "undelegation_queue_entry_list": undelegation_queue_entry_list,
      "queue_state_undelegation": copy.deepcopy(old["app_state"]["delegation"]["queue_state_undelegation"]),
      "redelegation_cooldown_list": []  # reset all cooldowns
    }

    # distribution -> unbond everything
    new["app_state"]["distribution"] = {
      "params": {
        "community_tax": "0.020000000000000000",
        "base_proposer_reward": "0.000000000000000000",
        "bonus_proposer_reward": "0.000000000000000000",
        "withdraw_addr_enabled": True
      },
      "fee_pool": {
        "community_pool": []
      },
      "delegator_withdraw_infos": copy.deepcopy(old["app_state"]["distribution"]["delegator_withdraw_infos"]),
      "previous_proposer": "",
      "outstanding_rewards": [],  # clear all outstanding rewards
      "validator_accumulated_commissions": [],  # clear all outstanding rewards
      "validator_historical_rewards": [],  # clear all outstanding rewards
      "validator_current_rewards": [],  # clear all outstanding rewards
      "delegator_starting_infos": [],  # clear all outstanding rewards
      "validator_slash_events": []  # clear all outstanding rewards
    }

    # pool -> migrate to v1.4 pools
    pool_list = list()
    for pool in copy.deepcopy(old["app_state"]["pool"]["pool_list"]):
        pool_list.append({
          'config': pool["config"],
          'current_compression_id':  pool["current_compression_id"],
          'current_index': pool["current_index"],
          'current_key': pool["current_key"],
          'current_storage_provider_id': pool["current_storage_provider_id"],
          'current_summary': pool["current_summary"],
          'disabled':  pool["disabled"],
          'id':  pool["id"],
          'logo':  pool["logo"],
          'max_bundle_size':  pool["max_bundle_size"],
          'min_delegation':  "100000000000",
          'name':  pool["name"],
          'inflation_share_weight': "1000000",
          'protocol':  pool["protocol"],
          'runtime': pool["runtime"],
          'start_key': pool["start_key"],
          'total_bundles': pool["total_bundles"],
          'upgrade_plan': pool["upgrade_plan"],
          'upload_interval': pool["upload_interval"]
        })

    new["app_state"]["pool"] = {
      "params": {
        "protocol_inflation_share": "0.010000000000000000",
        "pool_inflation_payout_rate": "0.050000000000000000"
      },
      "pool_list": pool_list,
      "pool_count": old["app_state"]["pool"]["pool_count"]
    }

    # funders -> create funding state for all pools
    new['app_state']['funders']['funding_state_list'] = \
        [{'pool_id': pool["id"], 'active_funder_addresses': []} for pool in old["app_state"]["pool"]["pool_list"]]

    # stakers -> partial reset
    converted_stakers = copy.deepcopy(old["app_state"]["stakers"]["staker_list"])
    for entry in converted_stakers:
        entry["commission_rewards"] = "0"

    new["app_state"]["stakers"] = {
      "params": {
        "commission_change_time": "3600",
        "leave_pool_time": "3600"
      },
      "staker_list": converted_stakers,
      "valaccount_list": [],  # reset valaccounts
      "commission_change_entries": [],  # reset commission changes
      "queue_state_commission": {  # reset commission changes
        "low_index": "0",
        "high_index": "0"
      },
      "leave_pool_entries": [],  # reset leave pool entries
      "queue_state_leave": {  # reset leave pool entries
        "low_index": "0",
        "high_index": "0"
      }
    }

    # staking -> full reset
    new["app_state"]["staking"] = {
      "params": {
        "unbonding_time": "1814400s",
        "max_validators": 100,
        "max_entries": 7,
        "historical_entries": 10000,
        "bond_denom": "tkyve",
        "min_commission_rate": "0.000000000000000000"
      },
      "last_total_power": "0",
      "last_validator_powers": [],
      "validators": [],
      "delegations": [],  # delegations are paid back in `migrateBalances`
      "unbonding_delegations": [],  # reset
      "redelegations": [],  # reset
      "exported": False
    }

    # gov -> keep gov counter and params
    new["app_state"]["gov"] = {
        "starting_proposal_id": old["app_state"]["gov"]["starting_proposal_id"],
        "deposits": [],  # reset
        "votes": [],  # reset
        "proposals": [],  # reset
        "deposit_params": None,
        "voting_params": None,
        "tally_params": None,
        "params": {
            "min_deposit": [
                {
                    "denom": "tkyve",
                    "amount": "25000000000"
                }
            ],
            "max_deposit_period": "600s",
            "voting_period": "600s",
            "quorum": "0.334000000000000000",
            "threshold": "0.500000000000000000",
            "veto_threshold": "0.334000000000000000",
            "min_initial_deposit_ratio": "0.100000000000000000",
            "burn_vote_quorum": False,
            "burn_proposal_deposit_prevote": False,
            "burn_vote_veto": True
        }
    }


# Needs to run after the normal migration has happened.
# Assumes only tkyve balances with 6 decimals
def migrateBalances(new_genesis, old_genesis):

    # move delegations back to delegators, ignore slashes
    delegation_returned = dict()
    for entry in old_genesis["app_state"]["staking"]["delegations"]:
        delegation_returned.setdefault(entry["delegator_address"], 0)
        delegation_returned[entry["delegator_address"]] += int(float(entry["shares"]) / 1000)

    for entry in new_genesis["app_state"]["bank"]["balances"]:
        if entry["address"] in delegation_returned:
            entry["coins"][0]["amount"] = str(int(entry["coins"][0]["amount"]) + delegation_returned[entry["address"]])

    print("Delegation Returned", sum(delegation_returned.values()))

    total_supply = 0
    for entry in new_genesis["app_state"]["bank"]["balances"]:
        if entry["coins"][0]["amount"] == "0":
            entry["coins"] = []
            continue

        if entry["address"] in [
            # burn gov deposits if there are any
            "kyve10d07y265gmmuvt4z0w9aw880jnsr700jdv7nah",
            # clear ecosystem
            "kyve1hfvhl7vf635xta2l4y5p4myj23pp7sg08f5rew",
            # clear team account
            "kyve10a48445a7vtdyce8f6rzxq9qsjtxqx6zeac50m",
            # clear distribution module
            "kyve1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8s3f2zm",
            # clear Round 1
            "kyve1cp4xc3zjmllnjkp9rkv86r7e472ca99jvcg4ln",
            # clear Round 2
            "kyve1zgerj6467dvu8hvgcv93d8v2ngmhajnat04ccq",
            # clear Round 3
            "kyve1mhe3xkysswtc37m8xmcfvt3jdw6qpkun5ts2ug",
            # clear ValidatorUS
            "kyve170zz90rvt8a8uj6t3869y80j67j8ln8ednwqqs",
            # clear Marketing
            "kyve1t7rc09jgzsxtkuw56c0nmj88nzuc0mwe3pe3tf",
            # clear ValidatorEU
            "kyve1rqcgwqau3zw5zehmzz29jkeckt7pqqk507sd8c",
            # clear bonded_tokens_pool
            "kyve1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3z4yejn",
            # clear KYVE registry
            "kyve1sujfrgcdvr2e393dumnmsd9tw6e25e0mgd3mc4",
            # clear Validator-2
            "kyve1dsnj3v7yh86ktwu4l299p4dqmepjjhukprqy7t",
            # clear Validator-1
            "kyve154faqxswn6stqde9vmhd7g59xrnyejlzszqduh",
            # clear not_bonded_tokens_pool
            "kyve1tygms3xhhs3yv487phx3dw4a95jn7t7lk4cgy8",
            # clear pool account
            "kyve1yl9v25pcxem9e5g828f84d9xu97h4qx58sw7yx",
        ]:
            entry["coins"] = []
            continue

        total_supply += int(entry["coins"][0]["amount"])

    new_genesis["app_state"]["bank"]["supply"] = [{'amount': str(total_supply), 'denom': 'tkyve'}]


def addGenesisBalances(genesis, address, amount, account_type="/cosmos.auth.v1beta1.BaseAccount"):
    for entry in genesis["app_state"]["auth"]["accounts"]:
        if entry["@type"] == "/cosmos.auth.v1beta1.BaseAccount":
            if entry["address"] == address:
                break
        elif entry["@type"] == "/cosmos.auth.v1beta1.ModuleAccount":
            if entry["base_account"]["address"] == address:
                break
        else:
            if entry["base_vesting_account"]["base_account"]["address"] == address:
                break
    else:
        genesis["app_state"]["auth"]["accounts"].append(
            {
                '@type': account_type,
                'address': address,
                'pub_key': None,
                'account_number': "0",
                'sequence': '0'
            }
        )

    for entry in genesis["app_state"]["bank"]["balances"]:
        if entry["address"] == address:
            if len(entry["coins"]) == 0:
                entry["coins"] = [{'amount': "0", 'denom': 'tkyve'}]

            entry["coins"][0]["amount"] = str(int(entry["coins"][0]["amount"]) + int(amount))
            break
    else:
        genesis["app_state"]["bank"]["balances"].append(
            {
                'address': address,
                'coins': [{'amount': str(amount), 'denom': 'tkyve'}]
            }
        )

    genesis["app_state"]["bank"]["supply"][0]["amount"] = (
        str(int(genesis["app_state"]["bank"]["supply"][0]["amount"])
            + int(amount))
    )


def injectValidators(genesis, validators):
    genesis["validators"] = list()
    for validator in validators:
        addGenesisBalances(genesis,validator["address"], validator["amount"])

        # Add to tendermint validators
        genesis["validators"].append(
            {
                "address": validator["private_validator_key_address"],
                "name": validator["name"],
                "power": str(int(validator["stake"] / 1000000)),
                "pub_key": {
                    "type": "tendermint/PubKeyEd25519",
                    "value": validator["pub_key"]
                }
            }
        )
        # default delegator_starting_infos
        genesis["app_state"]["distribution"]["delegator_starting_infos"].append(
            {
                "delegator_address": validator["address"],
                "starting_info": {
                    "height": str(genesis["initial_height"]),
                    "previous_period": "1",
                    "stake": str(validator["stake"]) + ".000000000000000000"
                },
                "validator_address": validator["valoper_address"]
            }
        )
        # default empty rewards
        genesis["app_state"]["distribution"]["outstanding_rewards"].append(
            {
                "outstanding_rewards": [],
                "validator_address": validator["valoper_address"]
            }
        )
        # default empty commissions
        genesis["app_state"]["distribution"]["validator_accumulated_commissions"].append(
            {
                "accumulated": {
                    "commission": []
                },
                "validator_address": validator["valoper_address"]
            }
        )
        # default empty rewards
        genesis["app_state"]["distribution"]["validator_current_rewards"].append(
            {
                "rewards": {
                    "period": "2",
                    "rewards": []
                },
                "validator_address": validator["valoper_address"]
            }
        )
        # default historical rewards
        genesis["app_state"]["distribution"]["validator_historical_rewards"].append(
            {
                "period": "1",
                "rewards": {
                    "cumulative_reward_ratio": [],
                    "reference_count": 2
                },
                "validator_address": validator["valoper_address"]
            }
        )
        # initial delegations
        genesis["app_state"]["staking"]["delegations"].append(
            {
                "delegator_address": validator["address"],
                "shares": str(validator["stake"]) + ".000000000000000000",
                "validator_address": validator["valoper_address"]
            }
        )
        # initial last validator powers
        genesis["app_state"]["staking"]["last_validator_powers"].append(
            {
                "address": validator["valoper_address"],
                "power": str(int(validator["stake"] / 1000000))
            }
        )

        # initialise validator profile
        genesis["app_state"]["staking"]["validators"].append(
            {
                "commission": {
                    "commission_rates": {
                        "max_change_rate": "0.010000000000000000",
                        "max_rate": "0.200000000000000000",
                        "rate": "0.100000000000000000"
                    },
                    "update_time": "2022-05-31T10:53:22.074136565Z"
                },
                "consensus_pubkey": {
                    "@type": "/cosmos.crypto.ed25519.PubKey",
                    "key": validator["pub_key"]
                },
                "delegator_shares": str(validator["stake"]) + ".000000000000000000",
                "description": {
                    "details": "",
                    "identity": "",
                    "moniker": validator["name"],
                    "security_contact": "",
                    "website": ""
                },
                "jailed": False,
                "min_self_delegation": "1",
                "operator_address": validator["valoper_address"],
                "status": "BOND_STATUS_BONDED",
                "tokens": str(validator["stake"]),
                "unbonding_height": "0",
                "unbonding_time": "1970-01-01T00:00:00Z"
            }
        )
        # default slashing
        genesis["app_state"]["slashing"]["missed_blocks"].append(
            {
                "address": validator["valcons_address"],
                "missed_blocks": []
            }
        )
        genesis["app_state"]["slashing"]["signing_infos"].append(
            {
                "address": validator["valcons_address"],
                "validator_signing_info": {
                    "address": validator["valcons_address"],
                    "index_offset": "1",
                    "jailed_until": "1970-01-01T00:00:00Z",
                    "missed_blocks_counter": "0",
                    "start_height": "0",
                    "tombstoned": False
                }
            }
        )

    # total voting power
    genesis["app_state"]["staking"]["last_total_power"] = str(
        int(genesis["app_state"]["staking"]["last_total_power"]) + sum([v["amount"] for v in validators]))

    # add stake to bonded tokens pool
    addGenesisBalances(genesis, "kyve1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3z4yejn",
                       sum([v["stake"] for v in validators]))


if __name__ == '__main__':
    if len(sys.argv[1:]) != 2:
        print("Usage: python3 KorelliaRegenesis.py [old_genesis.json] [output_genesis.json]")
        exit(1)

    print("Reading file ...")

    with open(sys.argv[1]) as genesisFile:
        old_korellia = json.load(genesisFile)

    # obtain configured v47 genesis
    new_genesis = initialiseDefaultGenesis()

    # migrate old Korellia data into the new genesis
    migrateKorelliaToNewGenesis(new_genesis, old_korellia)

    # cleanup Korellia balances
    migrateBalances(new_genesis, old_korellia)

    # setup new validators
    injectValidators(new_genesis, [
        {
            # Validator EU
            "address": "kyve1ghpmzfuggm7vcruyhfzrczl4aczy8gas4cus34",
            "valoper_address": "kyvevaloper1ghpmzfuggm7vcruyhfzrczl4aczy8gas8guslh",
            "valcons_address": "kyvevalcons1cm8ldtrvmycahclvp8hxyddv0cz556a0qm8d08",
            "private_validator_key_address": "C6CFF6AC6CD931DBE3EC09EE6235AC7E054A6BAF",
            "name": "Validator-EU",
            "amount": 100000000000000,
            "stake": 100000000000000,
            "pub_key": "PxDjqoMTb8rDsMLIfDKudGTyXWE4ruTjcGwFrE3JxGc="
        },
        {
            # Validator US
            "address": "kyve1p6ctexlukvllyruwnyfhh2cvdwqggz95kjqxh8",
            "valoper_address": "kyvevaloper1p6ctexlukvllyruwnyfhh2cvdwqggz95yzqxe9",
            "valcons_address": "kyvevalcons10gk2crjkhahq28x38kx2l77x4fq4h7204a3ee8",
            "private_validator_key_address": "7A2CAC0E56BF6E051CD13D8CAFFBC6AA415BF94F",
            "name": "Validator-US",
            "amount": 100000000000000,
            "stake": 100000000000000,
            "pub_key": "kA7gxEDoJ+459wZEv+A3ErZaI3cpIyzJ5wQtRWZPFLo="
        }
    ])

    # KYVE Ecosystem
    addGenesisBalances(new_genesis, "kyve1ygtqlzxhvp3t0wwcjd5lmq4zxl0qcck9g3mmgp", "400000000000000")

    # Add Team account
    addGenesisBalances(new_genesis, "kyve1e29j95xmsw3zmvtrk4st8e89z5n72v7nf70ma4", "100000000000000", "/cosmos.auth.v1beta1.ModuleAccount")

    print("Writing file ...")
    with open(sys.argv[2], 'w') as f:
        mergedJsonString = json.dumps(new_genesis)
        f.write(mergedJsonString)

    print("Done")
