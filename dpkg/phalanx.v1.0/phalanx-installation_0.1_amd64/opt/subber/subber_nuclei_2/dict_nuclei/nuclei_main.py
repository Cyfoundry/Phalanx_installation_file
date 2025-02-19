import nu_main
import json
import argparse

parser = argparse.ArgumentParser()
parser.add_argument("-t", "--target", help="input the target", dest="target", default="127.0.0.1", type=str)
parser.add_argument("-m", "--masterid", help="input the masterid", dest="masterid", default="", type=str)
args = parser.parse_args()

target = args.target
masterid = args.masterid

final_results = []  # 用於存儲所有目標的掃描結果
target_list = target.split(',')
total_score_sum = 0  # 各個IP分數加總
total_ip_count = len(target_list)  # 總IP數

for target in target_list:
    result = nu_main.main(target, masterid)
    if result:
        final_results.append(result)
        total_score_sum += result.get("score", 0)

# 根據提供的算式計算總分
if total_ip_count > 0:  # 確保分母不為零
    total_score = 100 - 100 * (total_score_sum / (total_ip_count * 10))
else:
    total_score = 100  # 如果沒有IP，則假定總分為100

# 將總分數包含在最終的JSON結果中
output = {
    "total_score": round(total_score, 2),
    "result": final_results
}

print(json.dumps(output, indent=2))

