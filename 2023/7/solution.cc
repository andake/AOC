#include <algorithm>
#include <cctype>
#include <charconv>
#include <cstddef>
#include <cstdio>
#include <exception>
#include <filesystem>
#include <format>
#include <fstream>
#include <iostream>
#include <iterator>
#include <list>
#include <ostream>
#include <ranges>
#include <regex>
#include <set>
#include <string>
#include <tuple>
#include <utility>
#include <vector>

namespace fs = std::filesystem;

std::vector<std::string> parseFile() {
  fs::path filePath("input.txt");

  if (!fs::exists(filePath)) {
    std::cerr << "file not found: " << filePath << std::endl;
    return {};
  }

  std::ifstream file(filePath);
  std::vector<std::string> lines;

  if (file.is_open()) {
    std::string buffer((std::istreambuf_iterator<char>(file)),
                       std::istreambuf_iterator<char>());

    auto range = buffer | std::views::split('\n') |
                 std::views::transform([](const auto &str) {
                   return std::string(str.begin(), str.end());
                 });
    lines.assign(range.begin(), range.end());
    file.close();
  }

  return lines;
}

const int FIVE = 7;
const int FOUR = 6;
const int HOUSE = 5;
const int THREE = 4;
const int TWO = 3;
const int ONE = 2;
const int HIGH = 1;

struct Hand {
  int multiplier;
  int type;
  std::vector<int> cards;

  auto operator<=>(const Hand &rhs) const {
    if (type != rhs.type) {
      return type <=> rhs.type;
    }
    for (size_t i = 0; i < cards.size(); i++) {
      if (cards.at(i) != rhs.cards.at(i)) {
        return cards.at(i) <=> rhs.cards.at(i);
      }
    }

    return type <=> rhs.type;
  }
};

void solution1() {
  auto lines = parseFile();

  std::vector<Hand> hands;

  std::map<std::string, int> cards{
      {"A", 14}, {"K", 13}, {"Q", 12}, {"J", 11}, {"T", 10}};

  for (const auto &line : lines) {
    Hand hand;
    auto split = line | std::views::split(' ') |
                 std::views::transform([](const auto &str) {
                   return std::string(str.begin(), str.end());
                 });

    for (const auto &w : split) {
      if (w.size() == 5) {
        for (const auto &c : w) {
          auto s = std::string{c};
          try {
            auto num = std::stoi(s);
            hand.cards.push_back(num);
          } catch (std::exception e) {
            hand.cards.push_back(cards.find(s)->second);
          }
        }
      } else {
        hand.multiplier = std::stoi(w);
      }
    }

    std::map<int, int> counts;
    for (const auto card : hand.cards) {
      auto c = counts.find(card);
      if (c != counts.end()) {
        c->second++;
      } else {
        counts.insert({card, 1});
      }
    }

    if (counts.size() == 5) {
      hand.type = HIGH;
    } else if (counts.size() == 4) {
      hand.type = ONE;
    } else if (counts.size() == 3) {
      auto threes = counts | std::views::values |
                    std::views::filter([](int x) { return x == 3; });
      if (threes.empty()) {
        hand.type = TWO;
      } else {
        hand.type = THREE;
      }
    } else if (counts.size() == 2) {
      auto fours = counts | std::views::values |
                   std::views::filter([](int x) { return x == 4; });
      if (fours.empty()) {
        hand.type = HOUSE;
      } else {
        hand.type = FOUR;
      }
    } else if (counts.size() == 1) {
      hand.type = FIVE;
    }

    hands.push_back(hand);
  }

  std::sort(hands.begin(), hands.end());

  int sum = 0;
  int rank = 1;
  for (const auto &h : hands) {
    sum += rank * h.multiplier;
    rank++;
  }

  std::cout << std::format("{}\n", sum);
}

void solution2() {
  auto lines = parseFile();

  std::vector<Hand> hands;

  std::map<std::string, int> cards{
      {"A", 14}, {"K", 13}, {"Q", 12}, {"J", 1}, {"T", 10}};

  for (const auto &line : lines) {
    Hand hand;
    auto split = line | std::views::split(' ') |
                 std::views::transform([](const auto &str) {
                   return std::string(str.begin(), str.end());
                 });

    for (const auto &w : split) {
      if (w.size() == 5) {
        for (const auto &c : w) {
          auto s = std::string{c};
          try {
            auto num = std::stoi(s);
            hand.cards.push_back(num);
          } catch (std::exception e) {
            hand.cards.push_back(cards.find(s)->second);
          }
        }
      } else {
        hand.multiplier = std::stoi(w);
      }
    }

    std::map<int, int> counts;
    for (const auto card : hand.cards) {
      auto c = counts.find(card);
      if (c != counts.end()) {
        c->second++;
      } else {
        counts.insert({card, 1});
      }
    }

    if (counts.contains(1) && counts.size() > 1) {
      auto jit = counts.find(1);
      auto nj = jit->second;
      counts.erase(jit);

      auto max = 0;
      for (auto v : counts | std::views::values) {
        if (v > max) {
          max = v;
        }
      }

      std::for_each(counts.begin(), counts.end(), [nj, max](auto &it) {
        if (it.second == max && max != 5) {
          it.second += nj;
        }
      });
    }

    if (counts.size() == 5) {
      hand.type = HIGH;
    } else if (counts.size() == 4) {
      hand.type = ONE;
    } else if (counts.size() == 3) {
      auto threes = counts | std::views::values |
                    std::views::filter([](int x) { return x == 3; });
      if (threes.empty()) {
        hand.type = TWO;
      } else {
        hand.type = THREE;
      }
    } else if (counts.size() == 2) {
      auto fours = counts | std::views::values |
                   std::views::filter([](int x) { return x == 4; });
      if (fours.empty()) {
        hand.type = HOUSE;
      } else {
        hand.type = FOUR;
      }
    } else if (counts.size() == 1) {
      hand.type = FIVE;
    }

    hands.push_back(hand);
  }

  std::sort(hands.begin(), hands.end());

  int sum = 0;
  int rank = 1;
  for (const auto &h : hands) {
    sum += rank * h.multiplier;
    rank++;
  }

  std::cout << std::format("{}\n", sum);
}

int main() {
  solution1();
  solution2();

  return 0;
}