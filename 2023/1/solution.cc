#include <algorithm>
#include <cctype>
#include <cstdio>
#include <filesystem>
#include <format>
#include <fstream>
#include <iostream>
#include <list>
#include <ostream>
#include <ranges>
#include <regex>
#include <string>

namespace fs = std::filesystem;

std::list<std::string> parseFile() {
  fs::path filePath("input.txt");

  if (!fs::exists(filePath)) {
    std::cerr << "file not found: " << filePath << std::endl;
    return {};
  }

  std::ifstream file(filePath);
  std::list<std::string> lines;

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

  auto sum = 0;
  for (const auto &l : lines) {
    auto first =
        std::ranges::find_if(l, [](const auto &c) { return std::isdigit(c); });
    auto second = std::ranges::find_if(
        l | std::views::reverse, [](const auto &c) { return std::isdigit(c); });
    auto totStr = std::string(1, *first) + std::string(1, *second);
    sum += std::atoi(totStr.c_str());
  }

  std::cout << sum << std::endl;
}

void solution2() {
  auto lines = parseFile();

  auto conv = [](const std::string &str) -> std::string {
    std::map<std::string, std::string> conv{
        {"one", "1"},   {"two", "2"},   {"three", "3"},
        {"four", "4"},  {"five", "5"},  {"six", "6"},
        {"seven", "7"}, {"eight", "8"}, {"nine", "9"}};

    auto it = conv.find(str);
    if (it != conv.end()) {
      return it->second;
    }

    return str;
  };

  auto sum = 0;
  auto i = 1;
  for (const auto &l : lines) {
    auto newStr = std::regex_replace(l, std::regex("oneight"), "oneeight");
    newStr = std::regex_replace(newStr, std::regex("twone"), "twoone");
    newStr = std::regex_replace(newStr, std::regex("eightwo"), "eighttwo");

    std::regex nums("[1-9]|one|two|three|four|five|six|seven|eight|nine");

    auto it = std::sregex_iterator(newStr.begin(), newStr.end(), nums);
    auto firstStr = (*it).str();
    std::sregex_iterator last;

    for (; it != std::sregex_iterator(); it++) {
      last = it;
    }

    auto totStr = conv(firstStr) + conv((*last).str());
    std::cout << std::format("{}: {} {} {}\n", i, firstStr, (*last).str(),
                             totStr);
    sum += std::atoi(totStr.c_str());
    i++;
  }

  std::cout << sum << std::endl;
}

int main() {
  solution1();
  solution2();

  return 0;
}
