<?php

require_once (dirname(__FILE__). '/../PDFView.php');

define('FPDF_FONTPATH', dirname(__FILE__) . '/../../fpdf/font/');
require_once (dirname(__FILE__).  '/../../fpdf/mbfpdf.php');

function s($str) {
    return mb_convert_encoding($str, "SJIS", "UTF-8");
}

function getTax($v, $taxRate) {
    return ceil($v * $taxRate / 100);
}

function removeTax($v, $taxRate) {
    return floor($v * 100 / (100 + $taxRate));
}

class TQPDFViewImpl implements PDFView {
    private $pdf;
    
    private $startX = 20;
    private $accColorR = 195;
    private $accColorG = 13;
    private $accColorB = 35;

    public function __construct() {
        $this->pdf = new MBFPDF('P', 'mm', 'A4');
        $this->pdf->AddMBFont(GOTHIC ,'SJIS');
        $this->pdf->AddMBFont(MINCHO ,'SJIS');
        $this->pdf->AddMBFont(HIRAKAKU_W3 ,'SJIS');        
        $this->pdf->AddMBFont(HIRAKAKU_W6 ,'SJIS');
        $this->pdf->SetAutoPageBreak(false);
        $this->pdf->SetMargins(0, 0, 0, 0);
        $this->pdf->AddPage('P');
    }

    public function writeTitle($title) {
        $this->pdf->SetFillColor($this->accColorR, $this->accColorG, $this->accColorB);
        $this->pdf->Rect($this->startX, 10.0, 1.0, 33.0, 'F');

        $this->pdf->SetFont(HIRAKAKU_W6, '', 14);
        $this->pdf->SetTextColor($this->accColorR, $this->accColorG, $this->accColorB);
        $this->pdf->SetXY(27, 38);
        $this->pdf->write(4, s($title));        
    }

    public function writeDate($number, $dateTime) {
        $this->pdf->SetFont(HIRAKAKU_W3,'', 6.5);
        $this->pdf->SetTextColor(77, 77, 77);
        $this->pdf->SetDrawColor(153, 153, 153);
        $this->pdf->SetXY(162, 10);
        $this->pdf->Cell(28.5, 5.8, s('No.'. $number), 'B', 2, 'R', 0);
        
        $this->pdf->SetXY(162, 15.8);
        $this->pdf->Cell(28.5, 5.8, s(date('Y年n月j日', $dateTime)), 'B', 2, 'R', 0);
    }

    public function writeCompany($name, $title) {
        $this->pdf->SetFont(HIRAKAKU_W6,'', 10);
        $this->pdf->SetTextColor(0, 0, 0);
        $this->pdf->SetDrawColor(51, 51, 51);
        $this->pdf->SetXY(19.4, 59);
        $this->pdf->Cell(81.8, 8.35, s($name. " ". $title), 'B', 2, 'L', 0);
    }

    public function writeMyCompany($info, $user) {
        $this->pdf->SetFont(HIRAKAKU_W6, '', 8);
        $this->pdf->SetXY(140.5, 54);
        $this->pdf->write(6, s($info['company_name']));
        
        $this->pdf->SetFont(HIRAKAKU_W3,'', 7);
        $this->pdf->SetXY(140.5, 60.5);
        $this->pdf->write(4.3, s('〒'. $info['company_zip']. ' '. $info['company_address']));
        $this->pdf->SetXY(140.5, 64.8);
        $this->pdf->write(4.3, s('担当：'. $user['display_name']. ' / '. $user['tel']));        
    }
    
    public function writeItemTable($x, $y, $items, $taxRate, $totalLabel) {
        $h1 = 8;    // 1行目
        $h2 = 3;      // 2行目
        $h3 = 2;      // 余白行
        $hBank = 12;  // 振込先用
        
        $w0 = $this->startX; // $w0は開始位置をあわせるために空白として生成
        $w1 = 70.5;
        $w2 = 29;
        $w3 = 20.9;
        $w4 = 20.9;
        $w5 = 28.5;
        
        // 余白を定義
        $mS = 2;
        $mM = 5;
        
        // 表のはじまる位置
        $startTableY = 78;
        
        // 見出し ----------------------------------------------------
        $this->pdf->SetFont(HIRAKAKU_W3,'', 8);
        $this->pdf->SetTextColor($this->accColorR, $this->accColorG, $this->accColorB);
        $this->pdf->SetDrawColor($this->accColorR, $this->accColorG, $this->accColorB);
        $this->pdf->SetXY(0, $startTableY);
        $this->pdf->Cell($w0, $h1, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w1, $h1, s('項目 / 明細'), 'B', 0, 'C', 0);
        $this->pdf->Cell($w2, $h1, s('単価'), 'B', 0, 'C', 0);
        $this->pdf->Cell($w3, $h1, s('数量'), 'B', 0, 'C', 0);
        $this->pdf->Cell($w4, $h1, s('単位'), 'B', 0, 'C', 0);
        $this->pdf->Cell($w5, $h1, s('金額'), 'B', 1, 'C', 0);

        $sum = 0;
        $tax = 0;
        $no = 1;
        for ($i = 0 ; $i < count($items) ; ++$i) {
            $item = $items[$i];
            $price = $item['unit_price'] * $item['amount'];
            if ($item['tax_type'] == 1) {
                $tax += getTax($price, $taxRate);
            } else if ($item['tax_type'] == 2) {
                $body = removeTax($price, $taxRate);
                $tax += ($price - $body);
            } 
            
            $sum += $price;

            if ($i == 0) {
                $this->pdf->SetFont(HIRAKAKU_W3,'', 8);
                $this->pdf->SetTextColor(0, 0, 0);
                $this->pdf->Cell($w0, $h1, s(''), 0, 0, 'C', 0);
                $this->pdf->Cell($w1, $h1, s($item['subject']), 0, 0, 'L', 0);
                $this->pdf->Cell($w2, $h1, s('¥'. number_format($item['unit_price'])), 0, 0, 'R', 0);
                $this->pdf->Cell($w3, $h1, s($item['amount']), 0, 0, 'C', 0);
                $this->pdf->Cell($w4, $h1, s($item['degree']), 0, 0, 'C', 0);
                $this->pdf->Cell($w5, $h1, s('¥'. number_format($price)), 0, 1, 'R', 0);
                
                $this->pdf->SetFont(HIRAKAKU_W3,'', 7);
                $this->pdf->SetTextColor(77, 77, 77);
                $this->pdf->Cell($w0, $h2, s(''), 0, 0, 'C', 0);
                $this->pdf->MultiCell($w1, $h2, s($item['memo']), 0, 1, 'L', 0);
                
                //空白行
                $this->pdf->SetDrawColor(153, 153, 153);
                $this->pdf->Cell($w0, $h3, s(''), 0, 0, 'C', 0);
                $this->pdf->Cell($w1+$w2+$w3+$w4+$w5, $h3, s(''), 'B', 1, 'L', 0);
            } else {
                $this->pdf->SetFont(HIRAKAKU_W3,'', 8);
                $this->pdf->SetTextColor(0, 0, 0);
                $this->pdf->Cell($w0, $h1, s(''), 0, 0, 'C', 0);
                $this->pdf->Cell($w1, $h1, s($item['subject']), 0, 0, 'L', 0);
                $this->pdf->Cell($w2, $h1, s('¥'. number_format($item['unit_price'])), 0, 0, 'R', 0);
                $this->pdf->Cell($w3, $h1, s($item['amount']), 0, 0, 'C', 0);
                $this->pdf->Cell($w4, $h1, s($item['degree']), 0, 0, 'C', 0);
                $this->pdf->Cell($w5, $h1, s('¥'. number_format($price)), 0, 1, 'R', 0);
                
                $this->pdf->SetFont(HIRAKAKU_W3,'', 7);
                $this->pdf->SetTextColor(77, 77, 77);
                $this->pdf->Cell($w0, $h2, s(''), 0, 0, 'C', 0);
                $this->pdf->MultiCell($w1, $h2, s($item['memo']), 0, 1, 'L', 0);
                
                //空白行
                $this->pdf->SetDrawColor(153, 153, 153);
                $this->pdf->Cell($w0, $h3, s(''), 0, 0, 'C', 0);
                $this->pdf->Cell($w1+$w2+$w3+$w4+$w5, $h3, s(''), 'B', 1, 'L', 0);                
            }
            ++$no;
        }
        $total = $sum + $tax;

        $this->pdf->Cell($w0+$w1+$w2+$w3+$w4+$w5, $h3, s(''), 0, 1, 'L', 0);

        $this->pdf->SetFont(HIRAKAKU_W6,'', 8);
        $this->pdf->SetTextColor(77, 77, 77);
        $this->pdf->Cell($w0, $h2+$h3, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w1+$w2+$w3+$w4, $h2+$h3, s('小計金額'), 0, 0, 'R', 0);
        $this->pdf->SetTextColor(0, 0, 0);
        $this->pdf->Cell($w5, $h2+$h3, s('¥'. number_format($sum)), 0, 1, 'R', 0);

        $this->pdf->SetTextColor(77, 77, 77);
        $this->pdf->Cell($w0, $h2+$h3, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w1+$w2+$w3+$w4, $h2+$h3, s('消費税（'. (int)$taxRate. '％）'), 0, 0, 'R', 0);
        $this->pdf->SetTextColor(0, 0, 0);
        $this->pdf->Cell($w5, $h2+$h3, s('¥'. number_format($tax)), 0, 1, 'R', 0);

        //空白行
        $this->pdf->Cell($w0+$w1+$w2+$w3+$w4+$w5, $h3, s(''), 0, 1, 'L', 0);

        // 合計 ----------------------------------------------------
        $this->pdf->SetFont(HIRAKAKU_W6,'', 10);
        $this->pdf->SetTextColor($this->accColorR, $this->accColorG, $this->accColorB);
        $this->pdf->SetDrawColor($this->accColorR, $this->accColorG, $this->accColorB);
        $this->pdf->Cell($w0, $h1+$h2, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w1+$w2+$w3+$w4, $h1+$h2, s($totalLabel), 'TB', 0, 'R', 0);
        $this->pdf->Cell($w5, $h1+$h2, s('¥'. number_format($total)), 'TB', 1, 'R', 0);

        return array('sum' => $sum, 'tax' => $tax, 'total' => $total);
    }

    public function writeTheTimeForQuotation($days) {
        $w0 = $this->startX; // $w0は開始位置をあわせるために空白として生成
        $w1 = 70.5;
        $w2 = 29;
        $w3 = 20.9;
        $w4 = 20.9;
        $w5 = 28.5;
        $h1 = 8;
        $h2 = 3;

        $this->pdf->SetFont(GOTHIC,'', 7);
        $this->pdf->SetTextColor(77, 77, 77);
        $this->pdf->Cell($w0, $h1, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w1+$w2+$w3+$w4+$w5, $h1, s('＃お見積期限：発行日より'. $days. '日間'), 0, 1, 'R', 0);

        $this->pdf->Cell($w0, $h1, s(''), 0, 1, 'L', 0);  // 行の余白用
    }
        
    public function writeTheTimeForPayment($time, $limitType) {
        $w0 = $this->startX; // $w0は開始位置をあわせるために空白として生成
        $w1 = 70.5;
        $w2 = 29;
        $w3 = 20.9;
        $w4 = 20.9;
        $w5 = 28.5;
        $h1 = 8;
        $h2 = 3;

        $date = new DateTime(null, new DateTimeZone('Asia/Tokyo'));
        $date->setTimestamp($time);
        $year = (int)$date->format('Y');
        $month = (int)$date->format('n');
        $month += (int)$limitType;
        if ($month > 12) {
            $year++;
            $month -= 12;
        }
        
        $this->pdf->SetFont(HIRAKAKU_W3,'', 7);
        $this->pdf->SetTextColor(77, 77, 77);
        $this->pdf->Cell($w0, $h1, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w1+$w2+$w3+$w4+$w5, $h1, s('＃お支払期限：'. $year. '年'. $month. '月末日'), 0, 1, 'R', 0);

        $this->pdf->Cell($w0, $h1, s(''), 0, 1, 'L', 0);  // 行の余白用
    }

    public function writeProduct($workFrom, $workTo, $product, $memo) {
        $w0 = $this->startX; // $w0は開始位置をあわせるために空白として生成
        $w1 = 70.5;
        $w2 = 29;
        $w3 = 20.9;
        $w4 = 20.9;
        $w5 = 28.5;
        $h1 = 8;
        $h2 = 3;
        $mS = 2;
        $mM = 5;
        
        $this->pdf->SetFont(HIRAKAKU_W3,'', 8);
        $this->pdf->SetDrawColor($this->accColorR, $this->accColorG, $this->accColorB);

        $this->pdf->SetTextColor($this->accColorR, $this->accColorG, $this->accColorB);
        $this->pdf->SetLineWidth(0.5);
        $this->pdf->Cell($w0, $h1, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w3, $h1, s('作業期間'), 'B', 1, 'L', 0);

        $this->pdf->Cell($w0, $h2, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w1+$w2+$w3+$w4+$w5, $h2, s(''), 0, 1, 'L', 0);
        $this->pdf->SetTextColor(0, 0, 0);
        $this->pdf->Cell($w0, $h2, s(''), 0, 0, 'C', 0);
        $this->pdf->MultiCell($w1+$w2+$w3+$w4+$w5, $h2, s(date('Y年n月j日', $workFrom). '〜'. date('Y年n月j日', $workTo)), 0, 'L', 0);
        $this->pdf->Cell($w0, $mS, s(''), 0, 1, 'L', 0);  // 行の余白用

        $this->pdf->SetTextColor($this->accColorR, $this->accColorG, $this->accColorB);
        $this->pdf->SetLineWidth(0.5);
        $this->pdf->Cell($w0, $h1, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w3, $h1, s('成果物'), 'B', 1, 'L', 0);
        //空白行
        $this->pdf->Cell($w0, $h2, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w1+$w2+$w3+$w4+$w5, $h2, s(''), 0, 1, 'L', 0);
        //内容
        $this->pdf->SetTextColor(0, 0, 0);
        $this->pdf->Cell($w0, $h2, s(''), 0, 0, 'C', 0);
        $this->pdf->MultiCell($w1+$w2+$w3+$w4+$w5, $h2, s($product), 0, 'L', 0);
    
        $this->pdf->Cell($w0, $mS, s(''), 0, 1, 'L', 0);  // 行の余白用
    }

    public function writeMemo($memo) {
        $w0 = $this->startX; // $w0は開始位置をあわせるために空白として生成
        $w1 = 70.5;
        $w2 = 29;
        $w3 = 20.9;
        $w4 = 20.9;
        $w5 = 28.5;
        $h1 = 8;
        $h2 = 3;
        $mS = 2;
        $mM = 5;
                
        // 備考 ----------------------------------------------------
        $this->pdf->SetTextColor($this->accColorR, $this->accColorG, $this->accColorB);
        $this->pdf->SetLineWidth(0.5);
        $this->pdf->Cell($w0, $h1, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w3, $h1, s('備考'), 'B', 1, 'L', 0);
        //空白行
        $this->pdf->Cell($w0, $h2, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w1+$w2+$w3+$w4+$w5, $h2, s(''), 0, 1, 'L', 0);
        //内容
        $this->pdf->SetTextColor(0, 0, 0);
        $this->pdf->Cell($w0, $h2, s(''), 0, 0, 'C', 0);
        $this->pdf->MultiCell($w1+$w2+$w3+$w4+$w5, $h2, s($memo), 0, 'L', 0);

        $this->pdf->Cell($w0, $mS, s(''), 0, 1, 'L', 0);  // 行の余白用

        // ひとこと ----------------------------------------------------
        $this->pdf->Cell($w0, $mM, s(''), 0, 1, 'L', 0);  // 行の余白用

        $this->pdf->Cell($w0, $h1, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w1+$w2+$w4+$w5, $h1, s('上記のとおり、お見積り申し上げます。'), 0, 1, 'L', 0);        
    }

    public function writeBankInfo($info) {
        $w0 = $this->startX; // $w0は開始位置をあわせるために空白として生成
        $w1 = 70.5;
        $w2 = 29;
        $w3 = 20.9;
        $w4 = 20.9;
        $w5 = 28.5;
        $h1 = 8;
        $hBank = 12;  // 振込先用

        $bankType = ($info['company_bank_type'] == '1') ? '普通' :
            (($info['company_bank_type'] == '2') ? '当座' : '定期');
        
        $this->pdf->Cell($w0, $h1, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w1+$w2+$w4+$w5, $h1, s('いつもお引き立ていただきありがとうございます。今後ともどうぞよろしくお願いいたします。'), 0, 1, 'L', 0);
        
        $this->pdf->SetTextColor($this->accColorR, $this->accColorG, $this->accColorB);
        $this->pdf->SetLineWidth(0.2);
        $this->pdf->Cell($w0, $hBank, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w3, $hBank, s('お振込先'), LTB, 0, 'C', 0);
        $this->pdf->SetTextColor(0, 0, 0);
        $this->pdf->Cell($w1+$w2+$w4+$w5, $hBank, s($info['company_bankname']. ' '. $bankType.
                                                    ' '. $info['company_bank_num']. '　'. $info['company_bank_name']), RTB, 0, 'L', 0);
        
        $this->pdf->Cell($w0, $mM, s(''), 0, 1, 'L', 0);  // 行の余白用        
    }    

    public function writeTotal($total) {
        //$this->pdf->SetFont(MINCHO,'', 16);
        //$this->pdf->SetXY(16, 100);
        //$this->pdf->write(4, s($total));
    }

    public function writeStamp() {
        $this->pdf->Image('stamp.png', 171, 34);
    }

    public function output($name) {
        $this->pdf->Output($name . '.pdf', 'D');
    }
}
?>