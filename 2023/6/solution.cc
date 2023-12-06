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

void solution1() {
  auto lines = parseFile();

  std::vector<int> times;
  for (const auto &word : std::views::split(lines[0], ' ')) {
    try {
      auto i = std::stoi(std::string(std::string_view(word)));
      times.push_back(i);
    } catch (std::exception e) {
    }
  }

  std::vector<int> dists;
  for (const auto &word : std::views::split(lines[1], ' ')) {
    try {
      auto i = std::stoi(std::string(std::string_view(word)));
      dists.push_back(i);
    } catch (std::exception e) {
    }
  }

  int winsProd = 1;
  for (size_t i = 0; i < times.size(); i++) {
    int wins = 0;
    auto t = times.at(i);
    auto d = dists.at(i);

    for (auto ped = 0; ped <= t; ped++) {
      auto tot = ped * (t - ped);
      if (tot > d) {
        wins++;
      }
    }
    if (wins > 0) {
      winsProd *= wins;
    }
  }

  std::cout << std::format("winsProd: {}\n", winsProd);
}

void solution2() {
  auto lines = parseFile();

  std::string time;
  for (const auto &word : std::views::split(lines[0], ' ')) {
    try {
      auto d = std::string(std::string_view(word));
      std::stoi(d);
      time.append(d);
    } catch (std::exception e) {
    }
  }

  std::string dist;
  for (const auto &word : std::views::split(lines[1], ' ')) {
    try {
      auto d = std::string(std::string_view(word));
      std::stoi(d);
      dist.append(d);
    } catch (std::exception e) {
    }
  }

  int wins = 0;
  auto t = std::stol(time);
  auto d = std::stol(dist);

  for (auto ped = 0; ped <= t; ped++) {
    auto tot = ped * (t - ped);
    if (tot > d) {
      wins++;
    } // else
  }

  std::cout << std::format("wins: {}\n", wins);
}

int main() {
  solution1();
  solution2();

  return 0;
}