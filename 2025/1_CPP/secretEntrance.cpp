#include <iostream>
#include <string>
#include <iostream>
#include <fstream>
using namespace std;

int main()
{
    int zeroHits = 0;
    int currentVal = 50;
    int max = 99;

    string inputText;

    ifstream MyReadFile("input.txt");

    while (getline(MyReadFile, inputText))
    {
        char direction = inputText[0];
        int count = stoi(inputText.substr(1));

        if (direction == 'R')
        {
            currentVal += count;
        }
        else
        {
            currentVal -= count;
        }

        // reduce the value to the range -99 to 99
        currentVal %= 100;
        if (currentVal < 0)
        {
            currentVal += 100; // if val is negative add 100 to bring it back to 0-99 range
        }

        if (currentVal == 0)
        {
            zeroHits++;
        }

        cout << currentVal << endl;
    }

    MyReadFile.close();

    cout << zeroHits << endl;
}