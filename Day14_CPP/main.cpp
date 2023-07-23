#include <algorithm>
#include <cstdlib>
#include <iostream>
#include <fstream>
#include <vector>
#include <string>

using namespace std;

// Common

struct Coord
{
    int x = 0;
    int y = 0;

    Coord operator+(const Coord& rhs);
    Coord operator-(const Coord& rhs);
    bool operator==(const Coord& rhs);
};
using Coords = vector<Coord>;
using Lines = vector<Coords>;
    
Coord Coord::operator+(const Coord& rhs)
{
    return {
        x + rhs.x,
        y + rhs.y
    };
}

Coord Coord::operator-(const Coord& rhs)
{
    return {
        x - rhs.x,
        y - rhs.y
    };
}

bool Coord::operator==(const Coord& rhs)
{
    return x == rhs.x && y == rhs.y;
}

enum Cell
{
    Air = '.',
    Rock = '#',
    Sand = 'O',
    SandSpawer = '+',
    Path = '~',
};
using Row = vector<Cell>;

struct Field
{
    int offset;
    size_t width;
    vector<Row> rows;
    size_t sandCount = 0;
};

// Input

Coords parseLine(string line);

Lines readInput(string filename)
{
    ifstream is;
    is.open(filename);

    Lines lines;
    string lineAsStr;
    while (std::getline(is, lineAsStr))
    {
        if (lineAsStr.size() == 0)
            continue;
        lines.push_back(
            parseLine(lineAsStr)
        ); 
    }
    return lines;
}

Coord parseCoord(string coordAsStr);

Coords parseLine(string line)
{
    size_t pos;
    Coords res;
    while ((pos = line.find(" -> ")) != string::npos)
    {
        res.push_back(parseCoord(line.substr(0, pos)));
        line = line.substr(pos + 4);
    }
    res.push_back(parseCoord(line));
    return res;
}

Coord parseCoord(string coordAsStr)
{
    auto sp = coordAsStr.find(",");
    return Coord {
        atoi(coordAsStr.substr(0, sp).c_str()),
        atoi(coordAsStr.substr(sp+1).c_str())
    };
}

// Processing

Coord cmpBy(Coord c1, Coord c2, int(*fn)(int, int))
{
    return {
        fn(c1.x, c2.x),
        fn(c1.y, c2.y)
    };
}

Coord cmpAllBy(Lines lines, int(*fn)(int, int))
{
    Coord res = lines.at(0).at(0);
    for (auto i : lines) {
        for (auto c: i) {
            res = cmpBy(res, c, fn);
        }
    }
    return res;
}

Coords interpolateLine(Coords line)
{
    Coords res;
    res.push_back(line.at(0));
    for (size_t i = 1; i < line.size(); i++) {
        Coord diff = line[i] - line[i-1];
        diff.x = min(max(diff.x, -1), 1);
        diff.y = min(max(diff.y, -1), 1);

        Coord tmp = line[i-1] + diff;
        while(!(tmp == line[i]))
        {
            res.push_back(tmp);
            tmp = tmp + diff;
        }
        res.push_back(line[i]);
    }
    return res;
}

void addRowToField(Field& f)
{
    f.rows.insert(
        f.rows.begin(),
        vector<Cell>(f.width, Air)
    );
}

Field buildField(Lines lines)
{
    Coord min = cmpAllBy(lines, [](int a, int b){ return std::min(a, b);});
    Coord max = cmpAllBy(lines, [](int a, int b){ return std::max(a, b);});
   
    // Create space to fall out
    min.x--;
    max.x++;

    Field field;
    field.offset = min.x;
    field.width = (max.x - min.x) + 1;
    auto floor = max.y + 2;
    for (int i = 0; i < floor; i++)
    {
        addRowToField(field);
    }

    for (auto stroke : lines) {
        auto pts = interpolateLine(stroke);
        for (auto pt : pts){
            pt.x -= field.offset;
            field.rows.at(pt.y).at(pt.x) = Rock;
        }
    }

    return field;
}

void printField(Field f)
{
    for (auto r : f.rows) {
        for (auto c : r) {
            cout << (char)c;
        }
        cout << "\n";
    }
    cout << endl;
}

bool spawnSandPart1(Field& f)
{
    Coord sand{500 - f.offset, 0};
    while (true)
    {
        if (sand.y+1 >= f.rows.size())
            return false;

        auto& nextRow = f.rows[sand.y+1];
        if (nextRow[sand.x] == Air) {
        } else if (nextRow[sand.x-1] == Air) {
            sand.x--;
        } else if (nextRow[sand.x+1] == Air) {
            sand.x++;
        } else {
            // Can't move
            break;
        }
        sand.y++;
    }

    f.rows[sand.y][sand.x] = Sand;
    if (sand.y == 0) {
        addRowToField(f);
    }
    f.sandCount++;
    return true;
}

void extendFieldLeft(Field& f)
{
    for (auto& row : f.rows) {
        row.insert(row.begin(), Air);
    }
    f.width++;
    f.offset--;
}

void extendFieldRight(Field& f)
{
    for (auto& row : f.rows) {
        row.push_back(Air);
    }
    f.width++;
}

bool spawnSandPart2(Field& f)
{
    Coord sand{500 - f.offset, 0};
    while (true)
    {
        if (sand.y+1 >= f.rows.size()) {
            break;
        }
        auto& nextRow = f.rows[sand.y+1];
        if (nextRow[sand.x] == Air) {
        } else if (nextRow.at(sand.x-1) == Air) {
            sand.x--;
        } else if (nextRow.at(sand.x+1) == Air) {
            sand.x++;
        } else {
            // Can't move
            break;
        }
        sand.y++;

        // check for left and right border
        if (sand.x == 0) {
            extendFieldLeft(f);
            sand.x++; // Because we changed the offset
        } else if (sand.x+1 >= f.width) {
            extendFieldRight(f);
        }
    }

    f.rows[sand.y][sand.x] = Sand;
    f.sandCount++;
    return sand.y != 0;
}
// Main

void runPart1(string file)
{
    auto input = readInput(file);
    auto field = buildField(input);
    printField(field);
    while (spawnSandPart1(field)) {}
    printField(field);
    cout << "Counted sand: " << field.sandCount << endl;
}

void runPart2(string file)
{
    auto input = readInput(file);
    auto field = buildField(input);
    printField(field);
    while (spawnSandPart2(field)) {}
    printField(field);
    cout << "Counted sand: " << field.sandCount << endl;
}

int main()
{
    // runPart1("test");
    // runPart1("input");
    // runPart2("test");
    runPart2("input");
}
