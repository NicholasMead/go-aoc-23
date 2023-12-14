package y23d13;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

class Main {
    public static void main(String[] args) throws IOException {
        var lavas = getLavas();

        var score = 0;

        for (var lava : lavas) {
            score += getScore(lava);
        }

        System.out.println(score);
    }

    private static int getScore(Lava lava) {
        var rows = 0;
        var columns = 0;

        for (var row : lava.mirrorRows()) {
            rows += row;
        }

        for (var column : lava.mirrorColumns()) {
            columns += column;
        }

        return rows * 100 + columns;
    }

    private static List<Lava> getLavas() throws IOException {
        var path = Path.of("input.txt");
        var input = Files.readAllLines(path);

        var buffer = new Lava();
        var lavas = new ArrayList<Lava>();

        for (var line : input) {
            if (line.isEmpty()) {
                lavas.add(buffer);
                buffer = new Lava();
            } else {
                buffer.add(line);
            }
        }
        lavas.add(buffer);

        return lavas;
    }
}

class Lava {
    private List<String> lines;

    public Lava() {
        this.lines = new ArrayList<String>();
    }

    public void add(String line) {
        this.lines.add(line);
    }

    public List<String> getLines() {
        return Collections.unmodifiableList(this.lines);
    }

    public List<Integer> mirrorRows() {
        return mirrors(this.lines);
    }

    public List<Integer> mirrorColumns() {
        return mirrors(transpose(this.lines));
    }

    private static List<String> transpose(List<String> lines) {
        var transposed = new ArrayList<String>();

        for (var x = 0; x < lines.get(0).length(); x++) {
            transposed.add("");
        }

        for (var x = 0; x < lines.get(0).length(); x++) {
            for (var y = 0; y < lines.size(); y++) {
                transposed.set(x, transposed.get(x) + lines.get(y).charAt(x));
            }
        }

        return transposed;
    }

    private static List<Integer> mirrors(List<String> lines) {
        List<Integer> mirrors = new ArrayList<>();
        for (int i = 1; i < lines.size(); i++) {

            var left = lines.subList(0, i);
            var right = lines.subList(i, lines.size());

            if (isMirror(left, right)) {
                mirrors.add(i);
            }
        }
        return mirrors;
    }

    private static Boolean isMirror(List<String> left, List<String> right) {
        var extent = Math.min(left.size(), right.size());
        var result = true;

        for (int i = 0; i < extent; i++) {
            var leftLine = left.get(left.size() - 1 - i);
            var rightLine = right.get(i);

            if (!leftLine.equals(rightLine)) {
                result = false;
                break;
            }
        }

        return result;
    }
}