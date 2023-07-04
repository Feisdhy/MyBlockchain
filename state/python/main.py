import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
from matplotlib import rcParams

config = {
    "font.family": 'serif',
    "font.size": 20,
    "mathtext.fontset": 'stix',
    "font.serif": ['SimSong'],
}
rcParams.update(config)


def get_data(csv_path, col):
    datas = []
    with open(csv_path) as rw:
        next(rw)
        next(rw)
        for line in rw:
            lines = line.replace("\n", "").split(",")
            if col == 0:
                datas.append(int(lines[col]))
            else:
                datas.append(float(lines[col]) / 1000000000)
            # print(lines)

    return datas


def draw_rw_rate():
    # accounts = get_data("output_2.csv", 0)
    times = get_data("../file/output_2.csv", 1)

    x = np.arange(0, 1000, 1)

    plt.figure(figsize=(7, 4.8), )

    plt.plot(x, times, '', color='#366DD8', markerfacecolor='none', label="Block STM", lw=2)
    plt.yticks(range(1, 12, 2), ['1', '3', '5', '7', '9', '11'], fontsize=24)
    plt.xticks(range(0, 1001, 200), ['0', '2', '4', '6', '8', '10'], fontsize=24)
    plt.xlabel("账户数量 (千万)", fontsize=26)
    plt.ylabel("MPT写入时延 (s)", fontsize=26)
    plt.ylim(1, 12)

    plt.subplots_adjust(bottom=0.2, left=0.15)
    plt.grid(axis='y')
    # ax = plt.gca()
    # ax.margins(0.01)
    # plt.tight_layout()
    # font = {'size': 18, 'weight': "normal"}
    # plt.legend(bbox_to_anchor=(1.035, 1.26), loc=1, frameon=False, prop=font, ncol=2,
    #            labelspacing=0.3, columnspacing=1, handletextpad=1, handlelength=2)
    plt.savefig("insert.eps", bbox_inches='tight', pad_inches=0.0)
    plt.show()


if __name__ == '__main__':
    draw_rw_rate()
    # get_data("../data/classic.csv")


