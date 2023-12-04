#include <algorithm>
#include <cctype>
#include <charconv>
#include <cstddef>
#include <cstdio>
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

void solution1() {
  auto lines = parseFile();

  int sum = 0;
  for (const auto &line : lines) {
    auto range = line | std::views::split(':') |
                 std::views::transform([](const auto &str) {
                   return std::string(str.begin(), str.end());
                 });
    auto r = std::ranges::begin(range);

    auto range2 = *++r | std::views::split('|') |
                  std::views::transform([](const auto &str) {
                    return std::string(str.begin(), str.end());
                  });
    auto r2 = std::ranges::begin(range2);

    auto range3 = *r2 | std::views::split(' ') |
                  std::views::transform([](const auto &str) {
                    return std::string(str.begin(), str.end());
                  });

    std::set<int> wcs;

    for (const std::string &cv : range3) {
      if (!cv.empty()) {
        wcs.insert(std::stoi(cv));
      }
    }

    auto range4 = *++r2 | std::views::split(' ') |
                  std::views::transform([](const auto &str) {
                    return std::string(str.begin(), str.end());
                  });

    auto first = true;
    auto points = 0;
    for (const std::string &cv : range4) {
      if (!cv.empty()) {
        if (wcs.contains(std::stoi(cv))) {
          if (first) {
            points = 1;
            first = false;
          } else {
            points *= 2;
          }
        }
      }
    }

    sum += points;
  }

  std::cout << std::format("Sum: {}\n", sum);
}

void solution2() {
  auto lines = parseFile();

  std::vector<int> cards;
  for (size_t i = 0; i < lines.size(); i++) {
    auto cardNo = i + 1;
    auto line = lines[i];
    auto range = line | std::views::split(':') |
                 std::views::transform([](const auto &str) {
                   return std::string(str.begin(), str.end());
                 });
    auto r = std::ranges::begin(range);

    auto range2 = *++r | std::views::split('|') |
                  std::views::transform([](const auto &str) {
                    return std::string(str.begin(), str.end());
                  });
    auto r2 = std::ranges::begin(range2);

    auto range3 = *r2 | std::views::split(' ') |
                  std::views::transform([](const auto &str) {
                    return std::string(str.begin(), str.end());
                  });

    std::set<int> wcs;

    for (const std::string &cv : range3) {
      if (!cv.empty()) {
        wcs.insert(std::stoi(cv));
      }
    }

    auto range4 = *++r2 | std::views::split(' ') |
                  std::views::transform([](const auto &str) {
                    return std::string(str.begin(), str.end());
                  });

    auto wins = 0;
    for (const std::string &cv : range4) {
      if (!cv.empty()) {
        auto num = std::stoi(cv);
        if (wcs.contains(num)) {
          wins++;
        }
      }
    }

    int numCardNo = 1;
    for (const auto &cn : cards) {
      if (cn == (int)cardNo) {
        numCardNo++;
      }
    }

    for (auto i = 0; i < wins; i++) {
      for (auto j = 0; j < numCardNo; j++) {
        cards.push_back((int)cardNo + i + 1);
      }
    }
  }

  std::cout << std::format("Sum: {}\n", cards.size() + lines.size());
}

int main() {
  solution1();
  solution2();

  return 0;
}