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

struct Node {
  std::string left;
  std::string right;
};

void solution1() {
  auto lines = parseFile();

  auto instructions = lines.front();
  lines.erase(lines.begin(), lines.begin() + 2);

  std::map<std::string, Node> nodes;

  for (const auto &line : lines) {
    auto nodeName = line.substr(0, 3);
    auto left = line.substr(line.find_first_of("(") + 1, 3);
    auto right = line.substr(line.find_last_of(" ") + 1, 3);

    auto node = Node{.left = left, .right = right};

    nodes.insert({nodeName, node});
  }

  std::string currentNodeName = "AAA";
  auto steps = 0;
  size_t currentInstruction = 0;

  for (;;) {
    auto currentNode = nodes.find(currentNodeName);

    auto action = instructions.at(currentInstruction);
    if (action == 'L') {
      currentNodeName = currentNode->second.left;
    } else {
      currentNodeName = currentNode->second.right;
    }

    currentInstruction++;
    if (currentInstruction == instructions.size()) {
      currentInstruction = 0;
    }
    steps++;
    if (currentNodeName == "ZZZ") {
      break;
    }
  }

  std::cout << steps << std::endl;
}

void solution2() {}

int main() {
  solution1();
  solution2();

  return 0;
}