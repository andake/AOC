#include <algorithm>
#include <cctype>
#include <charconv>
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
#include <string_view>
#include <utility>

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

int getMax(const std::string &l, const std::string &color) {
  std::regex colorReg("([0-9]+) " + color);
  auto numColor = 0;
  for (auto it = std::sregex_iterator(l.begin(), l.end(), colorReg);
       it != std::sregex_iterator(); it++) {

    auto thisNumColor = std::stoi((*it).str(1));
    if (thisNumColor > numColor) {
      numColor = thisNumColor;
    }
  }

  return numColor;
}

void solution1() {
  auto lines = parseFile();

  const int maxRed = 12;
  const int maxGreen = 13;
  const int maxBlue = 14;

  int sum = 0;
  for (const auto &l : lines) {
    std::regex game("Game ([0-9]+)");
    std::smatch matches;
    auto id = 0;
    if (std::regex_search(l, matches, game)) {
      id = std::stoi(matches[1]);
    }

    auto blue = getMax(l, "blue");
    auto green = getMax(l, "green");
    auto red = getMax(l, "red");

    if (blue <= maxBlue && green <= maxGreen && red <= maxRed) {
      sum += id;
    }
  }

  std::cout << sum << std::endl;
}

void solution2() {
  auto lines = parseFile();
  int sum = 0;
  for (const auto &l : lines) {
    auto blue = getMax(l, "blue");
    auto green = getMax(l, "green");
    auto red = getMax(l, "red");

    sum += blue * green * red;
  }

  std::cout << sum << std::endl;
}

int main() {
  solution1();
  solution2();

  return 0;
}
