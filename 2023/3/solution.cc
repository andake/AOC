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

char getPos(size_t x, size_t y, const std::vector<std::string> &lines) {
  try {
    auto c = lines.at(y).at(x);
    return c;
  } catch (std::out_of_range &e) {
    return '.';
  }
}

bool isSign(auto c) { return std::isdigit(c) == 0 && c != '.'; }

bool isConnected(size_t x, size_t y, const std::vector<std::string> &lines) {
  if (isSign(getPos(x, y + 1, lines))) {
    return true;
  }
  if (isSign(getPos(x, y - 1, lines))) {
    return true;
  }
  if (isSign(getPos(x + 1, y, lines))) {
    return true;
  }
  if (isSign(getPos(x - 1, y, lines))) {
    return true;
  }
  if (isSign(getPos(x + 1, y + 1, lines))) {
    return true;
  }
  if (isSign(getPos(x + 1, y - 1, lines))) {
    return true;
  }
  if (isSign(getPos(x - 1, y + 1, lines))) {
    return true;
  }
  if (isSign(getPos(x - 1, y - 1, lines))) {
    return true;
  }

  return false;
}

void solution1() {
  auto lines = parseFile();

  int sum = 0;
  for (size_t y = 0; y != lines.size(); y++) {
    auto line = lines[y];

    for (size_t x = 0; x < line.size();) {
      char focus = line[x];
      if (std::isdigit(focus) != 0) {
        size_t forward = 1;
        std::string group{focus};

        bool connected = isConnected(x, y, lines);
        bool cont = false;

        if (auto next = getPos(x + 1, y, lines); std::isdigit(next) != 0) {
          group += next;
          forward++;
          if (!connected) {
            connected = isConnected(x + 1, y, lines);
          }
          cont = true;
        }
        if (auto next = getPos(x + 2, y, lines);
            std::isdigit(next) != 0 && cont) {
          group += next;
          forward++;
          if (!connected) {
            connected = isConnected(x + 2, y, lines);
          }
        }
        x += forward;

        if (connected) {
          sum += std::stoi(group);
        }
      } else {
        x++;
      }
    }
  }

  std::cout << std::format("Sum: {}\n", sum);
}

bool isStar(auto c) { return c == '*'; }

std::tuple<bool, size_t, size_t>
isGearCandidate(size_t x, size_t y, const std::vector<std::string> &lines) {
  if (isStar(getPos(x, y + 1, lines))) {
    return std::tuple{true, x, y + 1};
  }
  if (isStar(getPos(x, y - 1, lines))) {
    return std::tuple{true, x, y - 1};
  }
  if (isStar(getPos(x + 1, y, lines))) {
    return std::tuple{true, x + 1, y};
  }
  if (isStar(getPos(x - 1, y, lines))) {
    return std::tuple{true, x - 1, y};
  }
  if (isStar(getPos(x + 1, y + 1, lines))) {
    return std::tuple{true, x + 1, y + 1};
  }
  if (isStar(getPos(x + 1, y - 1, lines))) {
    return std::tuple{true, x + 1, y - 1};
  }
  if (isStar(getPos(x - 1, y + 1, lines))) {
    return std::tuple{true, x - 1, y + 1};
  }
  if (isStar(getPos(x - 1, y - 1, lines))) {
    return std::tuple{true, x - 1, y - 1};
  }

  return std::tuple{false, 0, 0};
}

void solution2() {
  auto lines = parseFile();
  int sum = 0;
  std::map<std::pair<size_t, size_t>, std::vector<int>> gearCandidates;
  for (size_t y = 0; y != lines.size(); y++) {
    auto line = lines[y];

    for (size_t x = 0; x < line.size();) {
      char focus = line[x];
      if (std::isdigit(focus) != 0) {
        size_t forward = 1;
        std::string group{focus};

        auto [gearCandidate, gx, gy] = isGearCandidate(x, y, lines);
        bool cont = false;

        if (auto next = getPos(x + 1, y, lines); std::isdigit(next) != 0) {
          group += next;
          forward++;
          if (!gearCandidate) {
            std::tie(gearCandidate, gx, gy) = isGearCandidate(x + 1, y, lines);
          }
          cont = true;
        }
        if (auto next = getPos(x + 2, y, lines);
            std::isdigit(next) != 0 && cont) {
          group += next;
          forward++;
          if (!gearCandidate) {
            std::tie(gearCandidate, gx, gy) = isGearCandidate(x + 2, y, lines);
          }
        }
        x += forward;

        if (gearCandidate) {
          auto groupInt = std::stoi(group);
          auto it = gearCandidates.find(std::pair<size_t, size_t>(gx, gy));
          if (it != gearCandidates.end()) {
            it->second.push_back(groupInt);
          } else {
            gearCandidates.insert({std::pair<size_t, size_t>(gx, gy),
                                   std::vector<int>{groupInt}});
          }
        }
      } else {
        x++;
      }
    }
  }

  for (const auto &[p, vec] : gearCandidates) {
    if (vec.size() > 1) {
      sum += vec.at(0) * vec.at(1);
    }
  }

  std::cout << std::format("Sum: {}\n", sum);
}

int main() {
  solution1();
  solution2();

  return 0;
}
