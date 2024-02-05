import math
import time
import curses
import os
import random

SAND_NUM = 36
HALF_H = 6
IS_EVEN = False


def pre_n_2(n: int):
    """

    :param n:
    :return: closet n**2 that is smaller than n
    """
    k = int(math.sqrt(n))
    while k * k > n:
        k -= 1
    return k


def pre_pow_2(n: int):
    """
    :param n:
    :return: closest pow of 2 that is smaller than n
    """
    n -= 1
    n |= n >> 1
    n |= n >> 2
    n |= n >> 4
    n |= n >> 8
    n |= n >> 16
    return (n + 1) // 2


def next_pow_2(n: int):
    """
    :param n:
    :return: closest pow of 2 that is larger than n
    """
    n -= 1
    n |= n >> 1
    n |= n >> 2
    n |= n >> 4
    n |= n >> 8
    n |= n >> 16
    return n + 1


def get_arithmetic_n(n: int):
    return 2 * n - 1


def current_state(n: int):
    """
    return current filled level, and the remaining level's sand num.
    Using the Arithmetic Sequence Sum Formula, know for a arithmetic
    sequence 1, 3, ..., 2 * n - 1, the sum of pre-n sequence is n ** 2.
    Firstly get the filled level, and then subtract it can get the surface.
    :param n:
    :return:
    """
    # upper part
    upper_cur_level = pre_n_2(n)
    upper_surface_sand_num = n - upper_cur_level * upper_cur_level

    # lower part
    lower_sand = SAND_NUM - n
    lower_cur_level = 0
    count_lower_level = HALF_H
    while lower_sand > get_arithmetic_n(count_lower_level):
        lower_cur_level += 1
        count_lower_level -= 1
        lower_sand -= get_arithmetic_n(count_lower_level)
    lower_surface_sand_num = lower_sand
    return upper_cur_level, upper_surface_sand_num, lower_cur_level, lower_surface_sand_num


def render_time_glass():
    pass


def main():
    stdscr = curses.initscr()
    curses.noecho()  # 不输出- -
    curses.cbreak()  # 立刻读取:暂不清楚- -
    stdscr.keypad(1)  # 开启keypad
    stdscr.box()

    width = os.get_terminal_size().columns
    height = os.get_terminal_size().lines
    c_y = height // 2 - 1
    c_x = width // 2 - 10


    total_width = get_arithmetic_n(HALF_H + 2)
    n = SAND_NUM

    vertical_offset = 10
    horizen_offset = 20
    while n >= 0:
        upper_cur_level, upper_surface_sand_num, _, _ = current_state(n)
        # print(upper_cur_level, upper_surface_sand_num, n)


        # render upper part
        for each_level in range(HALF_H):
            cur_space_num = total_width - get_arithmetic_n(HALF_H - each_level)

            for space_i in range(cur_space_num // 2 - 1):
                stdscr.addstr(vertical_offset + each_level, horizen_offset + space_i, ' ')

            stdscr.addstr(vertical_offset + each_level, horizen_offset + cur_space_num // 2 - 1, '\\')

            if HALF_H - each_level == upper_cur_level + 1:
                # total slots this level: get_arithmetic_n(HALF_H - each_level) -> 2 * (HALF_H - each_level) - 1
                # sands num: upper_surface_sand_num
                # space num: get_arithmetic_n(HALF_H - each_level) - upper_surface_sand_num

                # 0, get_arithmetic_n(HALF_H - each_level) - upper_surface_sand_num
                # get_arithmetic_n(HALF_H - each_level) - upper_surface_sand_num, get_arithmetic_n(HALF_H - each_level)

                for sand_i in range(cur_space_num // 2, cur_space_num // 2 + get_arithmetic_n(HALF_H - each_level) - upper_surface_sand_num):
                    stdscr.addstr(vertical_offset + each_level, horizen_offset + sand_i, ' ')
                for sand_i in range(cur_space_num // 2 + get_arithmetic_n(HALF_H - each_level) - upper_surface_sand_num, cur_space_num // 2 + get_arithmetic_n(HALF_H - each_level)):
                    stdscr.addstr(vertical_offset + each_level, horizen_offset + sand_i, '*')
            elif HALF_H - each_level > upper_cur_level:
                for sand_i in range(cur_space_num // 2, cur_space_num // 2 + get_arithmetic_n(HALF_H - each_level)):
                    stdscr.addstr(vertical_offset + each_level, horizen_offset + sand_i, ' ')
            elif HALF_H - each_level < upper_cur_level:
                for sand_i in range(cur_space_num // 2, cur_space_num // 2 + get_arithmetic_n(HALF_H - each_level)):
                    stdscr.addstr(vertical_offset + each_level, horizen_offset + sand_i, '*')
            stdscr.addstr(vertical_offset + each_level, horizen_offset + cur_space_num // 2 + get_arithmetic_n(HALF_H - each_level), '/')

            for space_i in range(cur_space_num // 2 + get_arithmetic_n(HALF_H - each_level) + 1, cur_space_num + get_arithmetic_n(HALF_H - each_level)):
                stdscr.addstr(vertical_offset + each_level, horizen_offset + space_i, ' ')

        time.sleep(1)
        n -= 1
        stdscr.move(c_y, c_x)
        stdscr.refresh()
    curses.endwin()

        # render lower part

        # zh_ = '1234567890-qwertyuiopasdfghjklzxcvbnm,[;l,]/~!@#$%^&*()_+}"?{:><}"'';'
        # for linei in range(1, width - 1):
        #     for linej in range(1, height - 1):
        #         if linej == c_y:
        #             if linei <= 5 or linei + 6 >= width:
        #                 stdscr.addstr(linej, linei, '$')
        #             else:
        #                 stdscr.addstr(linej, c_x, time.strftime('%Y-%m-%d %H:%M:%S'), curses.A_BOLD)
        #         else:
        #             randominx = random.randint(0, len(zh_) - 1)
        #             stdscr.addstr(linej, linei, zh_[randominx])

        # time.sleep(1)
        # n -= 1

def test():
    print(pre_n_2(9))
    print(pre_n_2(8))
    print(pre_n_2(10))
    print(pre_n_2(24))
    print(pre_n_2(25))
    print(pre_n_2(26))


if __name__ == '__main__':
    test()
