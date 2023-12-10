#include <algorithm>
#include <cctype>
#include <charconv>
#include <cmath>
#include <cstddef>
#include <cstdio>
#include <exception>
#include <filesystem>
#include <format>
#include <fstream>
#include <iostream>
#include <iterator>
#include <list>
#include <numeric>
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

void doStuff(std::vector<std::vector<int>> &stages) {
  std::vector<int> nextStage;

  auto vec = stages.at(stages.size() - 1);
  for (size_t i = 0; i < vec.size() - 1; i++) {
    auto e1 = vec.at(i);
    auto e2 = vec.at(i + 1);

    auto diff = e2 - e1;
    nextStage.push_back(diff);
  }
  stages.push_back(nextStage);

  auto isZero = [](int x) { return x != 0; };
  if (auto allZero = std::find_if(nextStage.begin(), nextStage.end(), isZero);
      allZero == std::end(nextStage)) {
    return;
  }

  doStuff(stages);
}

void solution1() {
  auto lines = parseFile();

  int sum = 0;
  for (const auto &line : lines) {
    std::vector<std::vector<int>> stages;
    auto range = line | std::views::split(' ') |
                 std::views::transform([](const auto &str) {
                   return std::string(str.begin(), str.end());
                 });
    std::vector<int> values;
    for (const auto &c : range) {
      values.push_back(std::stoi(c));
    }
    stages.push_back(values);

    doStuff(stages);

    std::reverse(stages.begin(), stages.end());

    for (size_t i = 0; i < stages.size(); i++) {
      auto &currVec = stages.at(i);
      if (i == 0) {
        currVec.push_back(0);
      } else {
        auto prevVec = stages.at(i - 1);
        auto diff =
            currVec.at(currVec.size() - 1) + prevVec.at(prevVec.size() - 1);
        currVec.push_back(diff);
      }
    }

    auto lastStage = stages.at(stages.size() - 1);
    sum += lastStage.at(lastStage.size() - 1);
  }

  std::cout << std::format("{}\n", sum);
}

void solution2() {
  auto lines = parseFile();

  int sum = 0;
  for (const auto &line : lines) {
    std::vector<std::vector<int>> stages;
    auto range = line | std::views::split(' ') |
                 std::views::transform([](const auto &str) {
                   return std::string(str.begin(), str.end());
                 });
    std::vector<int> values;
    for (const auto &c : range) {
      values.push_back(std::stoi(c));
    }
    stages.push_back(values);

    doStuff(stages);

    std::reverse(stages.begin(), stages.end());

    for (size_t i = 0; i < stages.size(); i++) {
      auto &currVec = stages.at(i);
      std::reverse(currVec.begin(), currVec.end());
      if (i == 0) {
        currVec.push_back(0);
      } else {
        auto prevVec = stages.at(i - 1);
        auto diff =
            currVec.at(currVec.size() - 1) - prevVec.at(prevVec.size() - 1);
        currVec.push_back(diff);
      }
    }

    auto lastStage = stages.at(stages.size() - 1);
    sum += lastStage.at(lastStage.size() - 1);
  }

  std::cout << std::format("{}\n", sum);
}

int main() {
  solution1();
  solution2();

  return 0;
}