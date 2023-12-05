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

struct ConvertMap {
  long src;
  long dst;
  long len;
};

long convert(const std::vector<ConvertMap> &map, long elem) {
  for (const auto &me : map) {
    if (elem >= me.src && elem <= me.src + me.len) {
      return me.dst + (elem - me.src);
    }
  }

  return elem;
}

void solution1() {
  auto lines = parseFile();
  std::vector<ConvertMap> ss;
  std::vector<ConvertMap> sf;
  std::vector<ConvertMap> fw;
  std::vector<ConvertMap> wl;
  std::vector<ConvertMap> lt;
  std::vector<ConvertMap> th;
  std::vector<ConvertMap> hl;
  std::vector<ConvertMap> *currMap = nullptr;

  for (const auto &line : lines) {
    if (line.empty()) {
      currMap = nullptr;
    }

    if (currMap != nullptr) {
      auto map =
          ConvertMap{.src = std::stol(line.substr(line.find_first_of(' ') + 1,
                                                  line.find_last_of(' ') - 1)),
                     .dst = std::stol(line.substr(0, line.find_first_of(' '))),
                     .len = std::stol(
                         line.substr(line.find_last_of(' '), line.size() - 1))};
      currMap->push_back(map);
    }

    if (line.starts_with("seed-to-soil map")) {
      currMap = &ss;
    } else if (line.starts_with("soil-to-fertilizer map")) {
      currMap = &sf;
    } else if (line.starts_with("fertilizer-to-water map")) {
      currMap = &fw;
    } else if (line.starts_with("water-to-light map")) {
      currMap = &wl;
    } else if (line.starts_with("light-to-temperature map")) {
      currMap = &lt;
    } else if (line.starts_with("temperature-to-humidity map")) {
      currMap = &th;
    } else if (line.starts_with("humidity-to-location map")) {
      currMap = &hl;
    }
  }

  auto seeds =
      lines[0].substr(lines[0].find_first_of(':') + 2, lines[0].size() - 1) |
      std::views::split(' ') | std::views::transform([](const auto &str) {
        return std::string(str.begin(), str.end());
      });

  std::vector<long> locs;
  for (const auto &seed : seeds) {
    auto s = convert(ss, std::stol(seed));
    auto s2 = convert(sf, s);
    auto s3 = convert(fw, s2);
    auto s4 = convert(wl, s3);
    auto s5 = convert(lt, s4);
    auto s6 = convert(th, s5);
    auto s7 = convert(hl, s6);
    locs.push_back(s7);
  }

  auto min = std::ranges::min(locs);
  std::cout << std::format("minLoc: {}\n", min);
}

long convert2(const std::vector<ConvertMap> &map, long elem) {
  for (const auto &me : map) {
    if (elem >= me.src && elem <= me.src + me.len) {
      return me.dst + (elem - me.src);
    }
  }

  return elem;
}


int main() {
  solution1();
  // solution2();

  return 0;
}